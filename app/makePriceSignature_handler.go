package app

import (
	"fmt"
	"github.com/TotemFi/totem-bridge-offchain/app/dto"
	"github.com/TotemFi/totem-bridge-offchain/common/httpx"
	"github.com/TotemFi/totem-bridge-offchain/model"
	"github.com/TotemFi/totem-bridge-offchain/swap"
	"github.com/TotemFi/totem-bridge-offchain/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
	"net/http"
)

// makePriceSignatureHandler godoc
// @Summary makePriceSignatureHandler
// @Description The Price Signature API
// @Tags bridge
// @Accept json
// @Produce json
// @Param account query string true "Account"
// @Param base query string true "Base"
// @Param quote query string true "Quote"
// @Param amount query string true "Amount"
// @Param fee query string true "Fee"
// @Param exchange query string true "Exchange"
// @Param deadline query uint64 true "Deadline"
// @Success 200 {object} dto.PriceSignatureResponseDto
// @Failure 400 {object} dto.ErrorResponseDto
// @Failure 401 {object} dto.ErrorResponseDto
// @Failure 500 {object} dto.ErrorResponseDto
// @Router /bridge/price/signature [get]
func (app *App) makePriceSignatureHandler(w http.ResponseWriter, r *http.Request) {

	account := r.URL.Query().Get("account")
	if len(account) > 0 && common.IsHexAddress(account) {
		util.Logger.Info("received account param in makePriceSignatureHandler: ", account)
	} else {
		httpx.JSONResponse(w, ErrMsg{Msg: "account param should not be empty"}, http.StatusBadRequest)
		return
	}

	base := r.URL.Query().Get("base")
	if len(base) > 0 {
		util.Logger.Info("received base param in makePriceSignatureHandler: ", base)
	} else {
		httpx.JSONResponse(w, ErrMsg{Msg: "base param should not be empty"}, http.StatusBadRequest)
		return
	}

	quote := r.URL.Query().Get("quote")
	if len(quote) > 0 {
		util.Logger.Info("received quote param in makePriceSignatureHandler: ", quote)
	} else {
		httpx.JSONResponse(w, ErrMsg{Msg: "quote param should not be empty"}, http.StatusBadRequest)
		return
	}

	amount := new(big.Int)
	queryParam := r.URL.Query().Get("amount")
	if len(queryParam) > 0 {
		var ok bool
		if amount, ok = amount.SetString(queryParam, 10); !ok {
			httpx.JSONResponse(w, ErrMsg{Msg: "amount param invalid"}, http.StatusBadRequest)
			return
		}
		util.Logger.Info("received amount param in makePriceSignatureHandler: ", amount)
	} else {
		httpx.JSONResponse(w, ErrMsg{Msg: "amount param should not be empty"}, http.StatusBadRequest)
		return
	}

	fee := new(big.Int)
	queryParam = r.URL.Query().Get("fee")
	if len(queryParam) > 0 {
		var ok bool
		if fee, ok = fee.SetString(queryParam, 10); !ok {
			httpx.JSONResponse(w, ErrMsg{Msg: "fee param invalid"}, http.StatusBadRequest)
			return
		}
		util.Logger.Info("received fee param in makePriceSignatureHandler: ", fee)
	} else {
		httpx.JSONResponse(w, ErrMsg{Msg: "fee param should not be empty"}, http.StatusBadRequest)
		return
	}

	exchange := new(big.Int)
	queryParam = r.URL.Query().Get("exchange")
	if len(queryParam) > 0 {
		var ok bool
		if exchange, ok = exchange.SetString(queryParam, 10); !ok {
			httpx.JSONResponse(w, ErrMsg{Msg: "exchange param invalid"}, http.StatusBadRequest)
			return
		}
		util.Logger.Info("received exchange param in makePriceSignatureHandler: ", exchange)
	} else {
		httpx.JSONResponse(w, ErrMsg{Msg: "exchange param should not be empty"}, http.StatusBadRequest)
		return
	}

	deadline := new(big.Int)
	queryParam = r.URL.Query().Get("deadline")
	if len(queryParam) > 0 {
		var ok bool
		if deadline, ok = deadline.SetString(queryParam, 10); !ok {
			httpx.JSONResponse(w, ErrMsg{Msg: "deadline param invalid"}, http.StatusBadRequest)
			return
		}
		util.Logger.Info("received deadline param in makePriceSignatureHandler: ", deadline)
	} else {
		httpx.JSONResponse(w, ErrMsg{Msg: "deadline param should not be empty"}, http.StatusBadRequest)
		return
	}

	var swapTx model.Swap
	err := app.DB.Where("sponsor = ?",
		account).Order("updated_at desc").Limit(1).First(&swapTx).Error
	if err != nil && err.Error() != "record not found" {
		util.Logger.Error("internal server error", err)
		httpx.JSONResponse(w, ErrMsg{Msg: "Something Was Wrong"}, http.StatusInternalServerError)
		return
	}

	if swapTx.ID > 0 && swapTx.State != swap.SwapStateCompleted && swapTx.State != swap.SwapStateNone {
		util.Logger.Info("account in active swapping  . . .", account)
		httpx.JSONResponse(w, ErrMsg{Msg: "account has an active peg-in request"}, http.StatusBadRequest)
		return
	}

	value, err := app.MetisBridgeAgent.SwapNonce(nil, common.HexToAddress(account))
	if err != nil {
		util.Logger.Error("get MetisBridgeAgent.SwapNonce failed ", account, err)
		httpx.JSONResponse(w, ErrMsg{Msg: "Something Was Wrong"}, http.StatusInternalServerError)
		return
	}
	nonce := new(big.Int).Add(value, big.NewInt(1))

	signerData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Bridge": []apitypes.Type{
				{Name: "swapType", Type: "string"},
				{Name: "account", Type: "address"},
				{Name: "base", Type: "string"},
				{Name: "quote", Type: "string"},
				{Name: "fee", Type: "uint256"},
				{Name: "amount", Type: "uint256"},
				{Name: "exchange", Type: "uint256"},
				{Name: "deadline", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
			},
		},
		PrimaryType: "Bridge",
		Domain: apitypes.TypedDataDomain{
			Name:              "totem bridge",
			Version:           "1",
			ChainId:           (*math.HexOrDecimal256)(big.NewInt(app.MetisChainID)),
			VerifyingContract: app.cfg.ChainConfig.MTSSwapAgentAddr,
		},
		Message: apitypes.TypedDataMessage{
			"swapType": "peg-in",
			"account":  account,
			"base":     base,
			"quote":    quote,
			"fee":      (*math.HexOrDecimal256)(fee),
			"amount":   (*math.HexOrDecimal256)(amount),
			"exchange": (*math.HexOrDecimal256)(exchange),
			"deadline": (*math.HexOrDecimal256)(deadline),
			"nonce":    (*math.HexOrDecimal256)(nonce),
		},
	}

	typedDataHash, err := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	if err != nil {
		util.Logger.Error("signerData.HashStruct of signerData.Message failed ", signerData.Message, err)
		httpx.JSONResponse(w, ErrMsg{Msg: "Invalid Data"}, http.StatusBadRequest)
		return
	}
	//util.Logger.Info("typedDataHash: ", typedDataHash.String())

	domainSeparator, err := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err != nil {
		util.Logger.Error("signerData.HashStruct of signerData.Domain failed ", signerData.Message, err)
		httpx.JSONResponse(w, ErrMsg{Msg: "Something Was Wrong"}, http.StatusInternalServerError)
		return
	}
	//util.Logger.Info("domainSeparatorHash: ", domainSeparator.String())

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	challengeHash := crypto.Keccak256Hash(rawData)

	signature, err := crypto.Sign(challengeHash.Bytes(), app.MetisPrivateKey)
	if err != nil {
		util.Logger.Error("crypto.HexToECDSA failed", err)
		httpx.JSONResponse(w, ErrMsg{Msg: "Something Was Wrong"}, http.StatusInternalServerError)
		return
	}
	signature[64] = signature[64] + 27

	httpx.JSONResponse(w, dto.PriceSignatureResponseDto{
		BaseAgent: app.cfg.ChainConfig.MTSSwapAgentAddr,
		Nonce:     nonce.String(),
		Signer:    app.MetisAdminAddress.String(),
		Signature: hexutil.Encode(signature),
	}, http.StatusOK)
	return
}
