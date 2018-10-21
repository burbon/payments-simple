package ps

type PaymentSourceResource struct {
	ID             string                 `json:"id"`
	Type           string                 `json:"type"`
	OrganisationID string                 `json:"organisation_id"`
	Version        string                 `json:"organisation_id"`
	Attributes     map[string]interface{} `json:"attributes"`
}

type PaymentsSource struct {
	Data []PaymentSourceResource `json:"data"`
}
