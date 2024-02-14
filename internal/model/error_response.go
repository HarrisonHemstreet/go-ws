package model

type ErrorResponse struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"` // Optional, for client-side handling
}
