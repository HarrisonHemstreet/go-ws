package model

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"` // Optional, for client-side handling
	StatusCode   int    `json:"status_code"`
}
