package repository

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"poc/app"
	"poc/model"
	"poc/repository/cassandra"
	"poc/repository/impls"
)

type (
	IRepositoryFactory interface {
		CreateRepository(repositoryType model.RepositoryType, appContext app.IAppContext) (impls.IRepositoryImpl, error)
	}

	RepositoryFactory struct{}
)

func NewRepositoryFactory() IRepositoryFactory {
	return &RepositoryFactory{}
}

func (rf *RepositoryFactory) CreateRepository(repositoryType model.RepositoryType, appContext app.IAppContext) (impls.IRepositoryImpl, error) {
	log.WithFields(log.Fields{"repo": repositoryType}).Info("Initializing repository")
	switch repositoryType {
	case model.CASSANDRA_REPOSITORY_TYPE:
		return cassandra.NewCassandraRepository(appContext), nil
	default:
		err := fmt.Errorf("CreateRepository: unsupported repositoryType: [%s]", repositoryType)
		return nil, err
	}
}
