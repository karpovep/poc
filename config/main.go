package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"poc/model"
)

type (
	ServerConfig struct {
		Port  string   `yaml:"port" env:"SERVER_PORT" env-default:":50051"`
		Nodes []string `yaml:"nodes"`
	}

	NodeServerConfig struct {
		Port string `yaml:"port" env:"NODE_SERVER_PORT" env-default:":40041"`
	}

	CassandraConfig struct {
		Hosts    []string
		Keyspace string
	}

	RepositoryConfig struct {
		Type model.RepositoryType `yaml:"type" env:"REPOSITORY_TYPE" env-default:"cassandra"`
	}

	ClientConfig struct {
		ServerAddress string `yaml:"serverAddress" env:"SERVER_ADDRESS" env-default:"localhost:50051"`
	}

	CloudConfig struct {
		NodeId     string `yaml:"nodeId"`
		Server     ServerConfig
		NodeServer NodeServerConfig `yaml:"nodeServer"`
		Client     ClientConfig
		Repository RepositoryConfig `yaml:"repository"`
		Cassandra  CassandraConfig
	}
)

func Init(configFile string) *CloudConfig {
	var cfg CloudConfig
	err := cleanenv.ReadConfig(configFile, &cfg)
	if err != nil {
		log.Fatal("Error reading config", err.Error())
	}
	return &cfg
}
