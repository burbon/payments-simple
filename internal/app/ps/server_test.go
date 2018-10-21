package ps

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
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
