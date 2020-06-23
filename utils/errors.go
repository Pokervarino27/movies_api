package utils

import "net/http"

//APIError error api handler
type APIError interface {
	Status() int
	Message() string
	Error() string
}

//apiError struct for a message response from API
type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AnError  string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.AStatus
}

func (e *apiError) Message() string {
	return e.AMessage
}

func (e *apiError) Error() string {
	return e.AnError
}

//NotFoundAPIError return an error message when the element is not found
func NotFoundAPIError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}
