package dto

type PriceSignatureRequestDto struct {
	Account  string `json:"account"`
	Base     string `json:"base"`
	Quote    string `json:"quote"`
	Amount   string `json:"amount"`
	Fee      string `json:"fee"`
	Exchange string `json:"exchange"`
	Deadline uint64 `json:"deadline"`
}

type PriceSignatureResponseDto struct {
	BaseAgent string `json:"baseAgent"`
	Nonce     string `json:"nonce"`
	Signer    string `json:"signer"`
	Signature string `json:"signature"`
}

//type PriceSignatureDto struct {
//	Types       PriceSignatureType `json:"types"`
//	PrimaryType string             `json:"primaryType"`
//	Domain      struct {
//		Name              string `json:"name"`
//		Version           string `json:"version"`
//		ChainID           *big.Int `json:"chainId"`
//		VerifyingContract string `json:"verifyingContract"`
//	} `json:"domain"`
//	Message struct {
//		SwapType string `json:"swapType"`
//		Base     string `json:"base"`
//		Quote    string `json:"quote"`
//		Fee      uint64 `json:"fee"`
//		Amount   *big.Int `json:"amount"`
//		Exchange *big.Int `json:"exchange"`
//		Deadline *big.Int `json:"deadline"`
//		None     string `json:"none"`
//		Account  string `json:"account"`
//	} `json:"message"`
//}
//
//type PriceSignatureType struct {
//	EIP712Domain []struct {
//		Name string `json:"name"`
//		Type string `json:"type"`
//	} `json:"EIP712Domain"`
//	Bridge []struct {
//		Name string `json:"name"`
//		Type string `json:"type"`
//	} `json:"bridge"`
//}
