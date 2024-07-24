package apierror

type APIError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

func (E *APIError) String() string {
	return string(E.Code) + E.Message
}

func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}
