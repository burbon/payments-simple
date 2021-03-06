package ps

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Specification struct {
	Debug             bool   `envconfig:"debug" default:"false"`
	Port              int    `envconfig:"babel_service_port_http" default:"8080"`
	PaymentsSourceURL string `default:"http://mockbin.org/bin/41ca3269-d8c4-4063-9fd5-f306814ff03f"`
	CassandraCluster  string `envconfig:"cassandra_cluster" default:"127.0.0.1"`
	CassandraKeyspace string `envconfig:"cassandra_keyspace" default:"paymentssimple"`
}

var Config Specification

func LoadConfig() {
	err := envconfig.Process("ps", &Config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if Config.Debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf(Config.PaymentsSourceURL)
}
