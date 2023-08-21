package app

import (
	"encoding/json"
	"fmt"
	"github.com/TotemFi/totem-bridge-offchain/app/dto"
	"github.com/TotemFi/totem-bridge-offchain/model"
	"github.com/TotemFi/totem-bridge-offchain/util"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
)

// UpdateSwapPairHandler godoc
// @Summary UpdateSwapPairHandler
// @Description The Update Swap Pair API
// @Tags admin
// @Accept json
// @Produce json
// @Param erc20_addr formData string true "ERC20Addr"
// @Param available formData string true "Available"
// @Param lower_bound formData string true "LowerBound"
// @Param upper_bound formData string true "UpperBound"
// @Param icon_url formData string true "IconUrl"
// @Success 200 {object} dto.UpdateSwapPairResponseDto
// @Failure 400 {object} dto.ErrorResponseDto
// @Failure 401 {object} dto.ErrorResponseDto
// @Failure 500 {object} dto.ErrorResponseDto
// @Router /admin/update_swap_pair [put]
func (app *App) UpdateSwapPairHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := app.checkAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updateSwapPair dto.UpdateSwapPairRequestDto
	err = json.Unmarshal(reqBody, &updateSwapPair)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := updateCheck(&updateSwapPair); err != nil {
		http.Error(w, fmt.Sprintf("parameters is invalid, %v", err), http.StatusBadRequest)
		return
	}

	swapPair := model.SwapPair{}
	err = app.DB.Where("erc20_addr = ?", updateSwapPair.ERC20Addr).First(&swapPair).Error
	if err != nil {
		http.Error(w, fmt.Sprintf("swapPair %s is not found", updateSwapPair.ERC20Addr), http.StatusBadRequest)
		return
	}

	toUpdate := map[string]interface{}{
		"available": updateSwapPair.Available,
	}

	if updateSwapPair.LowerBound != "" {
		toUpdate["low_bound"] = updateSwapPair.LowerBound
	}
	if updateSwapPair.UpperBound != "" {
		toUpdate["upper_bound"] = updateSwapPair.UpperBound
	}
	if updateSwapPair.IconUrl != "" {
		toUpdate["icon_url"] = updateSwapPair.IconUrl
	}

	err = app.DB.Model(model.SwapPair{}).Where("erc20_addr = ?", updateSwapPair.ERC20Addr).Updates(toUpdate).Error
	if err != nil {
		http.Error(w, fmt.Sprintf("update swapPair error, err=%s", err.Error()), http.StatusInternalServerError)
		return
	}

	// get swapPair
	swapPair = model.SwapPair{}
	err = app.DB.Where("erc20_addr = ?", updateSwapPair.ERC20Addr).First(&swapPair).Error
	if err != nil {
		http.Error(w, fmt.Sprintf("swapPair %s is not found", updateSwapPair.ERC20Addr), http.StatusBadRequest)
		return
	}

	swapPairIns, err := app.swapEngine.GetSwapPairInstance(common.HexToAddress(updateSwapPair.ERC20Addr))
	// disable is only for frontend, do not affect backend
	// if we want to disable it in backend, set the low_bound and upper_bound to be zero
	if err != nil && updateSwapPair.Available {
		// add swapPair in swapper
		err = app.swapEngine.AddSwapPairInstance(&swapPair)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if swapPairIns != nil {
		app.swapEngine.UpdateSwapInstance(&swapPair)
	}

	jsonBytes, err := json.MarshalIndent(swapPair, "", "  ")
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

func updateCheck(update *dto.UpdateSwapPairRequestDto) error {
	if update.ERC20Addr == "" {
		return fmt.Errorf("bsc_token_contract_addr can't be empty")
	}
	if update.UpperBound != "" {
		if _, ok := big.NewInt(0).SetString(update.UpperBound, 10); !ok {
			return fmt.Errorf("invalid upperBound amount: %s", update.UpperBound)
		}
	}
	if update.LowerBound != "" {
		if _, ok := big.NewInt(0).SetString(update.LowerBound, 10); !ok {
			return fmt.Errorf("invalid lowerBound amount: %s", update.LowerBound)
		}
	}
	if len(update.IconUrl) > MaxIconUrlLength {
		return fmt.Errorf("icon length exceed limit")
	}
	return nil
}
