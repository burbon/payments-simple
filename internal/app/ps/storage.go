package ps

import (
	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"

	strfmt "github.com/go-openapi/strfmt"
	"payments-simple/internal/pkg/rest/models"
)

func QCreatePaymentsFromJSON(payment []byte) {
	log.Debugf("payment %s", payment)
	query := `INSERT INTO payments JSON ?`
	if err := session.Query(query, payment).Exec(); err != nil {
		log.Fatalf("failed creating payment: %s", err)
	}
}

func QGetPayments() []*models.Payment {
	var payments []*models.Payment
	m := map[string]interface{}{}

	query := `SELECT id, type, organisation_id, version, attributes FROM payments`
	iterable := session.Query(query).Iter()
	for iterable.MapScan(m) {
		payments = append(payments, &models.Payment{
			ID:             strfmt.UUID(m["id"].(gocql.UUID).String()),
			Type:           m["type"].(string),
			OrganisationID: strfmt.UUID(m["organisation_id"].(gocql.UUID).String()),
			Version:        int64(m["version"].(int)),
			//Attributes:     &models.Attributes{
			//	Amount: m["attributes"].
			//},
		})
		m = map[string]interface{}{}
	}

	return payments
}
