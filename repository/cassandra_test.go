package repository

import (
	"github.com/gocql/gocql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"poc/app/app_mock"
	"poc/bus"
	"poc/bus/bus_mock"
	"poc/config"
	"poc/model"
	"poc/protos/cloud"
	"testing"
)

func Test_ShouldConnectToCassandraClusterAndCreateTable(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inboundChannelName := "inbound.test"
	inboundRepoChan := make(bus.DataChannel)

	keyspace := "cloud"
	createTableQueryParams := &CreateTableQueryParams{
		Keyspace:   keyspace,
		Table:      "test_table",
		PrimaryKey: "id",
		Fields: []struct {
			Name string
			Type string
		}{
			{"id", "UUID"},
			{"text", "text"},
		},
	}

	cfg := &config.CloudConfig{
		Cassandra: config.CassandraConfig{
			Hosts:    []string{"localhost"},
			Keyspace: keyspace,
		},
	}

	mockEventBus := bus_mock.NewMockIEventBus(mockCtrl)
	mockEventBus.EXPECT().Subscribe(inboundChannelName, inboundRepoChan)

	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("eventBus").Return(mockEventBus)
	mockAppContext.EXPECT().Get("config").Return(cfg)
	mockAppContext.EXPECT().Get("inboundRepoChan").Return(inboundRepoChan)
	mockAppContext.EXPECT().Get(model.INBOUND_CHANNEL_NAME).Return(inboundChannelName)

	cassandraRepo := NewCassandraRepository(mockAppContext)
	cassandraRepo.Start()

	// When
	err := cassandraRepo.CreateTable(createTableQueryParams)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ShouldSaveInternalServerObjectAndFindITByTypeAndId(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inboundChannelName := "inbound.test"
	inboundRepoChan := make(bus.DataChannel)

	entity := &cloud.TestEntity{Name: "My Test Entity"}
	entityType := string(entity.ProtoReflect().Descriptor().FullName())
	serialized, err := proto.Marshal(entity)
	if err != nil {
		t.Fatal("could not serialize:", err)
	}
	cloudObj := &cloud.CloudObject{
		Id:     gocql.TimeUUID().String(),
		Entity: &anypb.Any{TypeUrl: entityType, Value: serialized},
	}
	internalServerObject := model.NewIsoFromCloudObject(cloudObj)

	cfg := &config.CloudConfig{
		Cassandra: config.CassandraConfig{
			Hosts:    []string{"localhost"},
			Keyspace: "cloud",
		},
	}

	mockEventBus := bus_mock.NewMockIEventBus(mockCtrl)
	mockEventBus.EXPECT().Subscribe(inboundChannelName, inboundRepoChan)

	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("eventBus").Return(mockEventBus)
	mockAppContext.EXPECT().Get("config").Return(cfg)
	mockAppContext.EXPECT().Get("inboundRepoChan").Return(inboundRepoChan)
	mockAppContext.EXPECT().Get(model.INBOUND_CHANNEL_NAME).Return(inboundChannelName)

	cassandraRepo := NewCassandraRepository(mockAppContext)
	cassandraRepo.Start()

	// When
	err = cassandraRepo.SaveInternalServerObject(internalServerObject)
	if err != nil {
		t.Fatal(err)
	}
	actualObj, err := cassandraRepo.FindByTypeAndId(entityType, cloudObj.Id)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	assert.Equal(t, cloudObj, actualObj.CloudObj, "encoded saved and retrieved objects are not identical")

	var actualEntity cloud.TestEntity
	if err := actualObj.CloudObj.Entity.UnmarshalTo(&actualEntity); err != nil {
		t.Fatalf("Could not unmarshal TestEntity from the field: %s", err)
	}
	assert.Equal(t, entity.Name, actualEntity.Name, "after decoding: Name is not identical in saved and retrieved entities")
}
