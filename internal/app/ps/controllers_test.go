package ps

import (
	"encoding/json"
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	clearStorage()

	storedPayments := QGetPayments()
	assert.Equal(t, 0, len(storedPayments))

	resp := request("GET", "/fetch")
	assert.Equal(t, http.StatusOK, resp.Code)

	var body map[string]string
	log.Debugf("body: %s", resp.Body.String())
	err := json.Unmarshal([]byte(resp.Body.String()), &body)
	assert.Nil(t, err)

	assert.Equal(t, "fetched 14", body["message"])

	storedPayments = QGetPayments()
	assert.Equal(t, 14, len(storedPayments))
}

func TestGet(t *testing.T) {
	clearStorage()

	storedPayments := QGetPayments()
	assert.Equal(t, 0, len(storedPayments))

	_ = request("GET", "/fetch")
	resp := request("GET", "/")
	assert.Equal(t, http.StatusOK, resp.Code)

	var body []map[string]interface{}
	log.Infof("body: %s", resp.Body.String())
	err := json.Unmarshal([]byte(resp.Body.String()), &body)
	assert.Nil(t, err)

	assert.Equal(t, 14, len(body))
	assert.Equal(t, body[0]["id"], "09fe827a-b3c2-4437-b999-6c0e780c0983")

	bodyAttributes := body[0]["attributes"]
	log.Infof("attributes: %s", bodyAttributes)
	assert.NotEqual(t, "", bodyAttributes)
	assert.Equal(t, "100.21", bodyAttributes.(map[string]interface{})["amount"].(string))
}
