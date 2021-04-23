package repository

import (
	"github.com/golang/mock/gomock"
	"poc/app/app_mock"
	"poc/bus/bus_mock"
	"poc/config"
	"testing"
)

func Test_ShouldConnectToCassandraClusterCreateKeyspaceAndTable(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	keyspace := "test"
	createKeyspaceQueryParams := &CreateKeyspaceQueryParams{
		Keyspace:          keyspace,
		ReplicationClass:  "SimpleStrategy",
		ReplicationFactor: 1,
	}
	createTableQueryParams := &CreateTableQueryParams{
		Keyspace:   keyspace,
		Table:      "some_test_table",
		PrimaryKey: "id",
		Fields: []struct {
			Name string
			Type string
		}{
			{
				"id", "UUID",
			},
			{
				"text", "text",
			},
		},
	}

	cfg := &config.CloudConfig{
		Cassandra: config.CassandraConfig{
			Hosts: []string{"localhost"},
		},
	}

	mockEventBus := bus_mock.NewMockIEventBus(mockCtrl)

	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("eventBus").Return(mockEventBus)
	mockAppContext.EXPECT().Get("config").Return(cfg)

	cassandraRepo := NewCassandraRepository(mockAppContext)
	cassandraRepo.Start()

	// When
	err := cassandraRepo.CreateKeyspace(createKeyspaceQueryParams)
	if err != nil {
		t.Fatal(err)
	}
	err = cassandraRepo.CreateTable(createTableQueryParams)
	if err != nil {
		t.Fatal(err)
	}
}
