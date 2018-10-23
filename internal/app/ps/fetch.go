package ps

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"payments-simple/internal/pkg/rest/models"
)

func FetchPaymentsFromSource() (cnt int, err error) {
	payments, err := fetchPaymentsFromSource()
	cnt = len(payments.Data)

	for _, p := range payments.Data {
		CreatePayment(p)
	}

	return
}

func fetchPaymentsFromSource() (payments PaymentsSource, err error) {
	resp, err := http.Get(Config.PaymentsSourceURL)
	if err != nil {
		log.Fatalf("failed connecting to source: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed reading source response: %s", err)
		return
	}

	err = json.Unmarshal(body, &payments)
	if err != nil {
		log.Fatalf("failed unmarshaling source payments: %s, %s", err, body)
		return
	}

	return
}

func CreatePayment(payment models.Payment) {
	var rawPayment []byte
	var err error
	if rawPayment, err = json.Marshal(payment); err != nil {
		log.Fatalf("failed marshaling: %s", err)
	}
	QCreatePaymentsFromJSON(rawPayment)
}
