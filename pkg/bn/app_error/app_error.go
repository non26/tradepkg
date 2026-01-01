package apperror

import "fmt"

type AppError struct {
	code       int
	message    string
	rawMessage string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.code, e.message)
}

// code = error code
// message = convert error message to human readable message
// rawMessage = message from the cause of the error
func NewAppError(code int, message string, rawMessage string) *AppError {
	return &AppError{code: code, message: message, rawMessage: rawMessage}
}

func (e *AppError) GetCode() int {
	return e.code
}

func (e *AppError) GetMessage() string {
	return e.message
}

func (e *AppError) GetRawMessage() string {
	return e.rawMessage
}
