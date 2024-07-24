package apierror

type APIError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}
