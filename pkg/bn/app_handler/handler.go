package apphandler

import "github.com/labstack/echo/v4"

type IHandler interface {
	Handler(c echo.Context) (response interface{}, httpStatus int, err error)
}
