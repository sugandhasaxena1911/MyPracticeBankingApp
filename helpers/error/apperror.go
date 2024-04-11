package error

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"errormsg"`
}

func NewBadRequestAppError(msg string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: msg}

}
func NewNotFoundAppError(msg string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: msg}

}
func NewInternalServerAppError(msg string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: msg}

}
func (apperror *AppError) Getmessage() *AppError {
	return &AppError{Message: apperror.Message}
}
