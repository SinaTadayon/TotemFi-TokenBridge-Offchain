package dto

type SuccessResponseDto struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

type ErrorResponseDto struct {
	Code    string   `json:"code"`
	Error   string   `json:"error"`
	Details []string `json:"details"`
}
