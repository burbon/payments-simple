package ps

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func setup() {
	os.Setenv("PS_CASSANDRA_KEYSPACE", "test_paymentssimple")

	LoadConfig()
	CreateSession()
	router = SetupRouter()
}

func shutdown() {
	CloseSession()
}

func request(method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
