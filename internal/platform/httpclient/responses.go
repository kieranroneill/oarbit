package httpclient

type ListResponse[T any] struct {
	Items []T `json:"items"`
}
