package ps

import (
	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
)

var session *gocql.Session

func CreateSession() {
	cluster := gocql.NewCluster(Config.CassandraCluster)
	cluster.Keyspace = Config.CassandraKeyspace

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	log.Debug("Cassandra OK")

	return
}

func CloseSession() {
	session.Close()
}
