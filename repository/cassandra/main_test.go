package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"poc/app/app_mock"
	"poc/config"
	"poc/model"
	"poc/protos/cloud"
	"poc/repository/cassandra/queries"
	"testing"
)

func Test_ShouldConnectToCassandraClusterAndCreateTable(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	keyspace := "cloud"
	createTableQueryParams := &queries.CreateTableQueryParams{
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
			Hosts:         []string{"localhost"},
			Keyspace:      keyspace,
			TemplatesRoot: "queries/templates",
		},
	}

	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("config").Return(cfg)

	cassandraRepo := NewCassandraRepository(mockAppContext)
	cassandraRepo.Start()
	defer cassandraRepo.Stop()

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
	internalServerObject.Metadata.InitialNodeId = "test-node-1"

	cfg := &config.CloudConfig{
		Cassandra: config.CassandraConfig{
			Hosts:         []string{"localhost"},
			Keyspace:      "cloud",
			TemplatesRoot: "queries/templates",
		},
	}

	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("config").Return(cfg)

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
	assert.Equal(t, internalServerObject.Metadata.InitialNodeId, actualObj.Metadata.InitialNodeId, "encoded saved and retrieved Metadata.InitialNodeId is not identical")
}
