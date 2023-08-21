package dto

type PeginStateResponseDto struct {
	Id       string `json:"id"`
	Base     string `json:"base"`
	Quote    string `json:"quote"`
	Amount   string `json:"amount"`
	Fee      string `json:"fee"`
	Exchange string `json:"exchange"`
	Status   string `json:"status"`
}
