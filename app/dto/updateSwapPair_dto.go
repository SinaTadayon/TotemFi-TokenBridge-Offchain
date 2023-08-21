package dto

type UpdateSwapPairRequestDto struct {
	ERC20Addr  string `json:"erc20_addr"`
	Available  bool   `json:"available"`
	LowerBound string `json:"lower_bound"`
	UpperBound string `json:"upper_bound"`
	IconUrl    string `json:"icon_url"`
}

type UpdateSwapPairResponseDto struct {
	Sponsor    string `json:"sponsor"`
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	Decimals   int    `json:"decimals"`
	BEP20Addr  string `json:"bep_20_addr"`
	ERC20Addr  string `json:"erc_20_addr"`
	Available  bool   `json:"available"`
	LowBound   string `json:"low_bound"`
	UpperBound string `json:"upper_bound"`
	IconUrl    string `json:"icon_url"`
	RecordHash string `json:"record_hash"`
}
