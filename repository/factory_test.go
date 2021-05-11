package repository

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"poc/app/app_mock"
	"poc/config"
	"poc/model"
	"poc/repository/cassandra"
	"testing"
)

func TestRepositoryFactory(t *testing.T) {
	cases := []struct {
		name string
		exec func(t *testing.T)
	}{
		{
			"Should create ICassandraRepository successfully",
			func(t *testing.T) {
				// Given
				mockCtrl := gomock.NewController(t)
				defer mockCtrl.Finish()
				repositoryType := model.CASSANDRA_REPOSITORY_TYPE

				cfg := &config.CloudConfig{
					Cassandra: config.CassandraConfig{
						TemplatesRoot: "cassandra/queries/templates",
					},
				}

				mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
				mockAppContext.EXPECT().Get("config").Return(cfg)
				repositoryFactory := NewRepositoryFactory()

				// When
				repository, err := repositoryFactory.CreateRepository(repositoryType, mockAppContext)

				// Then
				if err != nil {
					t.Fatal("error creating repository", err)
				}
				_, ok := repository.(cassandra.ICassandraRepository)
				assert.True(t, ok, "It is expected that ICassandraRepository is created")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.exec(t)
		})
	}
}
