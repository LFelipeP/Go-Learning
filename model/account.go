package model

type Account struct {
	ID     int     `json:"id_account"`
	TaxId  string  `json:"taxId"`
	Amount float64 `json:"amount"`
}
