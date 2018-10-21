package ps

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	resp := request("GET", "/fetch")
	assert.Equal(t, http.StatusOK, resp.Code)

	var body map[string]string
	err := json.Unmarshal([]byte(resp.Body.String()), &body)
	assert.Nil(t, err)

	assert.Equal(t, "fetched 14", body["message"])
}
