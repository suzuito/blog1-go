package bgin

// ResponseError ...
type ResponseError struct {
	Detail string `json:"detail"`
}

// NewResponseError ...
func NewResponseError(err error) *ResponseError {
	return &ResponseError{
		Detail: err.Error(),
	}
}
