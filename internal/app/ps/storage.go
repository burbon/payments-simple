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
		log.Warnf("failed creating payment: %s", err)
	}
}

func QGetPayments() []*models.Payment {
	var payments []*models.Payment
	m := map[string]interface{}{}

	query := `SELECT id, type, organisation_id, version, attributes FROM payments`
	iterable := session.Query(query).Iter()
	for iterable.MapScan(m) {
		amount := m["attributes"].(map[string]interface{})["amount"].(string)
		payments = append(payments, &models.Payment{
			ID:             strfmt.UUID(m["id"].(gocql.UUID).String()),
			Type:           m["type"].(string),
			OrganisationID: strfmt.UUID(m["organisation_id"].(gocql.UUID).String()),
			Version:        int64(m["version"].(int)),
			Attributes: &models.Attributes{
				Amount: &amount,
			},
		})
		m = map[string]interface{}{}
	}

	return payments
}

func QGetPayment(payment_id strfmt.UUID) *models.Payment {
	m := map[string]interface{}{}
	query := `SELECT id, type, organisation_id, version, attributes FROM payments WHERE id = ?`
	id, _ := gocql.ParseUUID(payment_id.String())
	if err := session.Query(query, id).Consistency(gocql.One).MapScan(m); err != nil {
		log.Warnf("failed selecting payment: %s", err)
	}

	amount := m["attributes"].(map[string]interface{})["amount"].(string)
	payment := models.Payment{
		ID:             strfmt.UUID(m["id"].(gocql.UUID).String()),
		Type:           m["type"].(string),
		OrganisationID: strfmt.UUID(m["organisation_id"].(gocql.UUID).String()),
		Version:        int64(m["version"].(int)),
		Attributes: &models.Attributes{
			Amount: &amount,
		}}

	return &payment
}
