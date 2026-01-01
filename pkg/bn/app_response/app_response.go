package appresponse

import "github.com/labstack/echo/v4"

// general success
var SuccessCode string = "0000"
var SuccessMsg string = "Success"

// general error
var FailCode string = "9999"
var FailMsg string = "Failed"

// invalid request error
var InvalidRequestErrorCode string = "9000"
var InvalidRequestErrorMessage string = "Invalid request"

// found poisition in historty error
var FoundPositionInHistoryErrorCode string = "9001"
var FoundPositionInHistoryErrorMessage string = "Found position in history"

// duplicate client id for opening position error
var DuplicateClientIdForOpeningPositionErrorCode string = "9002"
var DuplicateClientIdForOpeningPositionErrorMessage string = "Duplicate client id for opening position"

// not found opening position error
var NotFoundOpeningPositionErrorCode string = "9003"
var NotFoundOpeningPositionErrorMessage string = "Not found opening position"

// sub account not registered error
var SubAccountNotRegisteredErrorCode string = "9004"
var SubAccountNotRegisteredErrorMessage string = "Sub account not registered"

type AppResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewAppResponse(code string, message string, data interface{}) *AppResponse {
	res := &AppResponse{Code: code, Message: message, Data: data}
	return res
}

func (a *AppResponse) SendResponse(httpStatus int, c echo.Context) error {
	return c.JSON(httpStatus, a)
}
