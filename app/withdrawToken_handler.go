package app

import (
	"encoding/json"
	"fmt"
	"github.com/TotemFi/totem-bridge-offchain/app/dto"
	cmm "github.com/TotemFi/totem-bridge-offchain/common"
	"github.com/TotemFi/totem-bridge-offchain/util"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
	"strings"
)

// WithdrawToken godoc
// @Summary WithdrawToken
// @Description The With Draw Token API
// @Tags admin
// @Accept json
// @Produce json
// @Param chain formData string true "Chain"
// @Param token_addr formData string true "TokenAddr"
// @Param recipient formData string true "Recipient"
// @Param amount formData string true "Amount"
// @Success 200 {object} dto.WithdrawTokenResponseDto
// @Failure 400 {object} dto.ErrorResponseDto
// @Failure 401 {object} dto.ErrorResponseDto
// @Failure 500 {object} dto.ErrorResponseDto
// @Router /admin/withdraw_token [post]
func (app *App) WithdrawToken(w http.ResponseWriter, r *http.Request) {
	reqBody, err := app.checkAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var withdrawToken dto.WithdrawTokenRequestDto
	err = json.Unmarshal(reqBody, &withdrawToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = withdrawCheck(&withdrawToken); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	amount := big.NewInt(0)
	amount.SetString(withdrawToken.Amount, 10)

	var withdrawResp dto.WithdrawTokenResponseDto
	//withdrawResp.TxHash, err = app.swapEngine.WithdrawToken(withdrawToken.Chain,
	//	common.HexToAddress(withdrawToken.TokenAddr),
	//	common.HexToAddress(withdrawToken.Recipient), amount)
	//if err != nil {
	//	withdrawResp.ErrMsg = err.Error()
	//}

	jsonBytes, err := json.MarshalIndent(withdrawResp, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		util.Logger.Errorf("write response error, err=%s", err.Error())
	}
}

func withdrawCheck(withdraw *dto.WithdrawTokenRequestDto) error {
	if strings.ToUpper(withdraw.Chain) != cmm.ChainBSC && strings.ToUpper(withdraw.Chain) != cmm.ChainMTS {
		return fmt.Errorf("bsc_token_contract_addr can't be empty")
	}
	if !common.IsHexAddress(withdraw.TokenAddr) {
		return fmt.Errorf("token address is not a valid address")
	}
	if !common.IsHexAddress(withdraw.Recipient) {
		return fmt.Errorf("recipient is not a valid address")
	}
	_, ok := big.NewInt(0).SetString(withdraw.Amount, 10)
	if !ok {
		return fmt.Errorf("invalid input, expected big integer")
	}
	return nil
}
