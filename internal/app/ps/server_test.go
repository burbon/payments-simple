package ps

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	router *gin.Engine
)

func setup() {
	router = SetupRouter()
}

func request(method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestPing(t *testing.T) {
	resp := request("GET", "/ping")
	assert.Equal(t, http.StatusOK, resp.Code)

	var body map[string]string
	err := json.Unmarshal([]byte(resp.Body.String()), &body)
	assert.Nil(t, err)

	assert.Equal(t, "pong", body["message"])
}
