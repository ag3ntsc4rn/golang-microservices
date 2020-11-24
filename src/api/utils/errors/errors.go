package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AnError   string `json:"error,omitempty"`
}

func (a *apiError) Status() int {
	return a.AStatus
}

func (a *apiError) Message() string {
	return a.AMessage
}

func (a *apiError) Error() string {
	return a.AnError
}

func NewApiErrorFromBytes(body []byte) (ApiError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("Invalid json body")
	}
	return &result, nil
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		AStatus: statusCode,
		AMessage: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
	}
}

func NewNotFoundError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}
