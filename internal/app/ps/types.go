package ps

import (
	"payments-simple/internal/pkg/rest/models"
)

type PaymentsSource struct {
	Data []models.Payment `json:"data"`
}
