package ps

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Specification struct {
	Debug             bool   `envconfig:"debug" default:"true"`
	PaymentsSourceURL string `default:"http://mockbin.org/bin/41ca3269-d8c4-4063-9fd5-f306814ff03f"`
}

var config Specification

func LoadConfig() {
	err := envconfig.Process("ps", &config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf(config.PaymentsSourceURL)
}
