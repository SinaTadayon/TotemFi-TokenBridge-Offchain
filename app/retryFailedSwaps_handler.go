package app

import (
	"encoding/json"
	"github.com/TotemFi/totem-bridge-offchain/app/dto"
	"github.com/TotemFi/totem-bridge-offchain/util"
	"net/http"
)

// RetryFailedSwaps godoc
// @Summary RetryFailedSwaps
// @Description The Retry Failed Swaps API
// @Tags admin
// @Accept json
// @Produce json
// @Param swap_id_list formData string true "SwapIDList"
// @Success 200 {object} dto.RetryFailedSwapsResponseDto
// @Failure 400 {object} dto.ErrorResponseDto
// @Failure 401 {object} dto.ErrorResponseDto
// @Failure 500 {object} dto.ErrorResponseDto
// @Router /admin/retry_failed_swaps [post]
func (app *App) RetryFailedSwaps(w http.ResponseWriter, r *http.Request) {
	reqBody, err := app.checkAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var retryFailedSwaps dto.RetryFailedSwapsRequestDto
	err = json.Unmarshal(reqBody, &retryFailedSwaps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var retryFailedSwapsResp dto.RetryFailedSwapsResponseDto
	//retryFailedSwapsResp.SwapIDList, retryFailedSwapsResp.RejectedSwapIDList, err = app.swapEngine.InsertRetryFailedSwaps(retryFailedSwaps.SwapIDList)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	retryFailedSwapsResp.ErrMsg = err.Error()
	//} else {
	//	w.WriteHeader(http.StatusOK)
	//}

	jsonBytes, err := json.MarshalIndent(retryFailedSwapsResp, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonBytes)
	if err != nil {
		util.Logger.Errorf("write response error, err=%s", err.Error())
	}
}
