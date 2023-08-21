package dto

type PriceRequestDto struct {
	Amount  string `json:"amount"`
	Base    string `json:"base"`
	Quote   string `json:"quote"`
	Account string `json:"account"`
}

type PriceResponseDto struct {
	Fee      string `json:"fee"`
	Exchange string `json:"exchange"`
	Amount   string `json:"amount"`
	Base     string `json:"base"`
	Quote    string `json:"quote"`
	Deadline uint64 `json:"deadline"`
}
