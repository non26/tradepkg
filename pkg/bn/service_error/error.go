package serviceerror

import "fmt"

type IError interface {
	GetCode() string
	Error() string
}

type ServiceError struct {
	Code    string
	Message string
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}

func NewServiceErrorWith(code string, err error) IError {
	return &ServiceError{Code: code, Message: err.Error()}
}

func (e *ServiceError) GetCode() string {
	return e.Code
}
