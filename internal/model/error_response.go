package model

import "time"

type ErrorResponse struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorMessage string    `json:"error_message"`
	ErrorCode    string    `json:"error_code"`
	Path         string    `json:"path"`
	StatusCode   int       `json:"status_code"`
}
