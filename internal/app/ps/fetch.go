package ps

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
)

func FetchPaymentsFromSource() (payments PaymentsSource, err error) {
	resp, err := http.Get(config.PaymentsSourceURL)
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

	CreatePayment(body)

	return
}

func CreatePayment(payment []byte) {
	payment = []byte(`{}`)
	//if err := session.Query(`INSERT INTO payments JSON '?'`, payment).Exec(); err != nil {
	//	log.Fatal(err)
	//}
	if err := session.Query(`INSERT INTO payments (id) values (?)`,
		gocql.TimeUUID()).Exec(); err != nil {
		log.Fatal(err)
	}
}
