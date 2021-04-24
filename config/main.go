package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type (
	ServerConfig struct {
		Port  string   `yaml:"port" env:"PORT" env-default:":50051"`
		Nodes []string `yaml:"nodes"`
	}

	CassandraConfig struct {
		Hosts    []string
		Keyspace string
	}

	ClientConfig struct {
		ServerAddress string `yaml:"serverAddress" env:"SERVER_ADDRESS" env-default:"localhost:50051"`
	}

	CloudConfig struct {
		Server    ServerConfig
		Client    ClientConfig
		Cassandra CassandraConfig
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
