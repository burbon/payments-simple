package ps

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// fixtures
var payment_id = "09fe827a-b3c2-4437-b999-6c0e780c0983"
var payments_no = 14

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

	assert.Equal(t, fmt.Sprintf("fetched %d", payments_no), body["message"])

	storedPayments = QGetPayments()
	assert.Equal(t, payments_no, len(storedPayments))
}

func TestClearLoad(t *testing.T) {
	clearStorage()

	storedPayments := QGetPayments()
	assert.Equal(t, 0, len(storedPayments))

	loadFixtures()
	storedPayments = QGetPayments()
	assert.Equal(t, payments_no, len(storedPayments))
}

func TestGet(t *testing.T) {
	clearStorage()
	loadFixtures()

	resp := request("GET", "/")
	assert.Equal(t, http.StatusOK, resp.Code)

	var body []map[string]interface{}
	log.Debugf("body: %s", resp.Body.String())
	err := json.Unmarshal([]byte(resp.Body.String()), &body)
	assert.Nil(t, err)

	assert.Equal(t, payments_no, len(body))
	assertPayment(t, body[0])
}

func TestGetSingle(t *testing.T) {
	clearStorage()
	loadFixtures()

	resp := request("GET", fmt.Sprintf("/%s", payment_id))
	assert.Equal(t, http.StatusOK, resp.Code)

	var body map[string]interface{}
	err := json.Unmarshal([]byte(resp.Body.String()), &body)
	assert.Nil(t, err)
	assertPayment(t, body)
}

func assertPayment(t *testing.T, payment map[string]interface{}) {
	assert.Equal(t, payment["id"], payment_id)

	bodyAttributes := payment["attributes"]
	log.Debugf("attributes: %s", bodyAttributes)
	assert.NotEqual(t, "", bodyAttributes)
	assert.Equal(t, "100.21", bodyAttributes.(map[string]interface{})["amount"].(string))
}
