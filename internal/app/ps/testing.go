package ps

import (
	"net/http"
	"net/http/httptest"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	handler http.Handler
)

func setup() {
	os.Setenv("PS_CASSANDRA_KEYSPACE", "test_paymentssimple")

	LoadConfig()
	CreateSession()
	handler = SetupHandler()
}

func shutdown() {
	CloseSession()
}

func request(method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w
}

func clearStorage() {
	query := `TRUNCATE payments`
	if err := session.Query(query).Exec(); err != nil {
		log.Warnf("failed truncating: %s", err)
	}
}

func loadFixtures() {
	_, _ = FetchPaymentsFromSource()
}
