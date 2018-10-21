package ps

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func setup() {
	LoadConfig()
	router = SetupRouter()
}

func request(method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
