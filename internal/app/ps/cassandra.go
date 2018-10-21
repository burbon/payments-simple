package ps

import (
	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
)

func CreateSession() (session *gocql.Session) {
	cluster := gocql.NewCluster(config.CassandraCluster)
	cluster.Keyspace = config.CassandraKeyspace

	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	log.Debug("Cassandra OK")

	return
}
