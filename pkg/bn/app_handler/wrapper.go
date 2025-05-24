package apphandler

import (
	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func HandlerWrapper(handler IHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		commonResponse := appresponse.NewCommonResponse()
		data, httpStatus, err := handler.Handler(c)
		if err != nil {
			if data != nil {
				commonResponse.FailWithData(err.Error(), data)
			} else {
				commonResponse.Fail(err.Error())
			}
			return c.JSON(httpStatus, commonResponse)
		}
		commonResponse.Success(data)
		return c.JSON(httpStatus, commonResponse)
	}
}
