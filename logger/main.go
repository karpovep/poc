package logger

import (
	log "github.com/sirupsen/logrus"
	"poc/config"
)

func Init(cfg config.LoggerConfig) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Only log the specified severity or above.
	logLvl, err := log.ParseLevel(cfg.Level)
	if err != nil {
		log.Fatalln("invalid logger level", cfg.Level)
	}
	log.SetLevel(logLvl)
}
