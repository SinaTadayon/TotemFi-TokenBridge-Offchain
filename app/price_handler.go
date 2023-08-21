package app

import (
	"fmt"
	"github.com/TotemFi/totem-bridge-offchain/app/dto"
	"github.com/TotemFi/totem-bridge-offchain/common/httpx"
	"github.com/TotemFi/totem-bridge-offchain/util"
	"math/big"
	"net/http"
	"time"
)

// makePriceHandler godoc
// @Summary makePriceHandler
// @Description The Price API Specifications
// @Tags bridge
// @Accept json
// @Produce json
// @Param base query string true "Base"
// @Param quote query string true "Quote"
// @Param amount query string true "Amount"
// @Param account query string true "Account"
// @Success 200 {object} dto.PriceResponseDto
// @Failure 400 {object} dto.ErrorResponseDto
// @Failure 401 {object} dto.ErrorResponseDto
// @Failure 500 {object} dto.ErrorResponseDto
// @Router /bridge/price [get]
func (app *App) makePriceHandler(w http.ResponseWriter, r *http.Request) {

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
		httpx.JSONResponse(w, ErrMsg{"amount param should not be empty"}, http.StatusBadRequest)
		return
	}

	if app.MetisLowerBoundAmount.Cmp(amount) > 0 || app.MetisUpperBoundAmount.Cmp(amount) < 0 {
		httpx.JSONResponse(w, ErrMsg{Msg: fmt.Sprintf("amount param must between %s and %s is valid",
			app.MetisLowerBoundAmount.String(), app.MetisUpperBoundAmount.String())}, http.StatusBadRequest)
		return
	}

	account := r.URL.Query().Get("account")
	if len(account) > 0 {
		util.Logger.Info("received account param in makePriceHandler: ", account)
	} else {
		httpx.JSONResponse(w, ErrMsg{Msg: "account param should not be empty"}, http.StatusBadRequest)
		return
	}

	//var swapTx model.Swap
	//err := app.DB.Where("sponsor = ?", account).Order("updated_at desc").Limit(1).First(&swapTx).Error
	//if err != nil && err.Error() != "record not found" {
	//	util.Logger.Error("account in active swapping not found", err)
	//	httpx.JSONResponse(w, "account not found", http.StatusNotFound)
	//	return
	//}
	//
	//if swapTx.ID > 0 && swapTx.State != swap.SwapStateCompleted && swapTx.State != swap.SwapStateNone {
	//	util.Logger.Info("account in active swapping  . . .", account)
	//	httpx.JSONResponse(w, "account has an active peg-in request", http.StatusBadRequest)
	//	return
	//}

	priceModel, err := app.priceUtils.SwapPriceCalc(amount)
	if err != nil {
		if err.Error() == "totem exchange is negative" {
			httpx.JSONResponse(w, ErrMsg{"Totem Exchange Is Negative, Metis Amount Invalid"}, http.StatusBadRequest)
			return
		} else {
			httpx.JSONResponse(w, ErrMsg{"Something Was Wrong"}, http.StatusInternalServerError)
			return
		}
	}

	deadline := time.Now().Add(10 * time.Minute).Unix()
	httpx.JSONResponse(w, dto.PriceResponseDto{
		Fee:      priceModel.TOTMFee.String(),
		Exchange: priceModel.TOTMExchange.String(),
		Amount:   amount.String(),
		Base:     base,
		Quote:    quote,
		Deadline: uint64(deadline),
	}, http.StatusOK)
	return
}
