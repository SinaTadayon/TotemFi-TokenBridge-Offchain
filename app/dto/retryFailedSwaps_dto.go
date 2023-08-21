package dto

type RetryFailedSwapsRequestDto struct {
	SwapIDList []uint `json:"swap_id_list"`
}

type RetryFailedSwapsResponseDto struct {
	SwapIDList         []uint `json:"swap_id_list"`
	RejectedSwapIDList []uint `json:"rejected_swap_id_list"`
	ErrMsg             string `json:"err_msg"`
}
