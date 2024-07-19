package model

type Transfer struct {
	ID                   int     `json:"id_transfer"`
	OriginAccountId      float64 `json:"originAccountId"`
	DestinationAccountId float64 `json:"destinationAccountId"`
	Amount               float64 `json:"amount"`
}
