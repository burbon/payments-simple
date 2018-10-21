package ps

type PaymentAttributes struct {
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	Reference string `json:"reference"`
}

type PaymentSourceResource struct {
	ID             string             `json:"id"`
	Type           string             `json:"type"`
	OrganisationID string             `json:"organisation_id"`
	Version        int                `json:"version"`
	Attributes     *PaymentAttributes `json:"attributes"`
}

type PaymentsSource struct {
	Data []PaymentSourceResource `json:"data"`
}
