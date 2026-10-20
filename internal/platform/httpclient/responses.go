package httpclient

type ListResponse[T any] struct {
	Items  []T   `json:"items"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"`
}
