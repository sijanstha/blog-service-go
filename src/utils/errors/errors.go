package errors

import (
	"errors"
	"net/http"

	dateutils "github.com/blog-service/src/utils/date"
)

type RestErr struct {
	Message   string `json:"message"`
	Code      int    `json:"code"`
	Error     string `json:"error"`
	Timestamp string `json:"timestamp"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message:   message,
		Code:      http.StatusBadRequest,
		Error:     "bad_request",
		Timestamp: dateutils.GetTodayDateInString(),
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message:   message,
		Code:      http.StatusNotFound,
		Error:     "not_found",
		Timestamp: dateutils.GetTodayDateInString(),
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message:   message,
		Code:      http.StatusInternalServerError,
		Error:     "internal_server_error",
		Timestamp: dateutils.GetTodayDateInString(),
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message:   message,
		Code:      http.StatusUnauthorized,
		Error:     "unauthorized_user",
		Timestamp: dateutils.GetTodayDateInString(),
	}
}

func NewError(msg string) error {
	return errors.New(msg)
}
