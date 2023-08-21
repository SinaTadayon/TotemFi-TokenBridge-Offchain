package app

import (
	"github.com/TotemFi/totem-bridge-offchain/app/dto"
	"github.com/TotemFi/totem-bridge-offchain/common"
	"github.com/TotemFi/totem-bridge-offchain/common/httpx"
	"github.com/TotemFi/totem-bridge-offchain/model"
	"github.com/TotemFi/totem-bridge-offchain/swap"
	"github.com/TotemFi/totem-bridge-offchain/util"
	"gorm.io/gorm/clause"
	"net/http"
)

// peginStateHandler godoc
// @Summary peginStateHandler
// @Description The Price Signature API
// @Tags bridge
// @Accept json
// @Produce json
// @Param account query string true "Account"
// @Param network query string true "Network"
// @Success 200 {object} dto.PeginStateResponseDto
// @Failure 400 {object} dto.ErrorResponseDto
// @Failure 401 {object} dto.ErrorResponseDto
// @Failure 500 {object} dto.ErrorResponseDto
// @Router /bridge/pegin/state [get]
func (app *App) peginStateHandler(w http.ResponseWriter, r *http.Request) {

	// if multiples possible, or to process empty values like param1 in
	// ?param1=&param2=something
	account := r.URL.Query().Get("account")
	if len(account) > 0 {
		util.Logger.Info("received account param in peginStateHandler: ", account)
	} else {
		httpx.JSONResponse(w, ErrMsg{Msg: "account param should not be empty"}, http.StatusBadRequest)
		return
	}

	network := r.URL.Query().Get("network")
	if len(network) > 0 {
		util.Logger.Info("received network param in peginStateHandler: ", network)
		// ... process them ... or you could just iterate over them without a check
		// this way you can also tell if they passed in the parameter as the empty string
		// it will be an element of the array that is the empty string
	} else {
		httpx.JSONResponse(w, ErrMsg{Msg: "network param should not be empty"}, http.StatusBadRequest)
		return
	}

	var swapTx model.Swap
	var data dto.PeginStateResponseDto
	err := app.DB.Preload(clause.Associations).Where("state in (?) and sponsor = ?", []common.SwapState{swap.SwapStateClaiming, swap.SwapStateSwapping, swap.SwapStateFilling}, account).Order("updated_at asc").Limit(1).First(&swapTx).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = app.DB.Preload(clause.Associations).Where("state in (?) and sponsor = ?", []common.SwapState{swap.SwapStateCompleted}, account).Order("updated_at desc").Limit(1).First(&swapTx).Error
			if err != nil {
				if err.Error() == "record not found" {
					data = dto.PeginStateResponseDto{
						Id:       "",
						Base:     "",
						Quote:    "",
						Amount:   "",
						Fee:      "",
						Exchange: "",
						Status:   string(swap.SwapStateNone),
					}
					httpx.JSONResponse(w, data, http.StatusOK)
					return
				} else {
					util.Logger.Error("internal server error", err)
					httpx.JSONResponse(w, ErrMsg{Msg: "Something Was Wrong"}, http.StatusInternalServerError)
					return
				}
			}
		} else {
			util.Logger.Error("internal server error", err)
			httpx.JSONResponse(w, ErrMsg{Msg: "Something Was Wrong"}, http.StatusInternalServerError)
			return
		}
	}

	data = dto.PeginStateResponseDto{
		Id:       swapTx.DataHash,
		Base:     swapTx.Base,
		Quote:    swapTx.Quote,
		Amount:   swapTx.SwapPrice.MetisAmount,
		Fee:      swapTx.SwapPrice.TOTMFee,
		Exchange: swapTx.SwapPrice.TOTMExchange,
		Status:   string(swapTx.State),
	}

	httpx.JSONResponse(w, data, http.StatusOK)
	return
}
