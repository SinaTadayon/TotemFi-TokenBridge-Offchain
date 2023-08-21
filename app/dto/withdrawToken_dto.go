package dto

type WithdrawTokenRequestDto struct {
	Chain     string `json:"chain"`
	TokenAddr string `json:"token_addr"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

type WithdrawTokenResponseDto struct {
	TxHash string `json:"tx_hash"`
	ErrMsg string `json:"err_msg"`
}
