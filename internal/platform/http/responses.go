package http

import "github.com/kieranroneill/oarbit/internal/platform/application"

//==========================================
// error response
//==========================================

type ErrorResponse struct {
	Errors []application.OarbitError `json:"errors"`
}

//==========================================
// list response
//==========================================

type ListResponse[T any] struct {
	Items  []T   `json:"items"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"`
}
