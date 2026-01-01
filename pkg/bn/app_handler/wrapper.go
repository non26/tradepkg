package apphandler

// func HandlerWrapper(handler IHandler) echo.HandlerFunc {
// return func(c echo.Context) error {
// 	commonResponse := appresponse.NewCommonResponse()
// 	data, httpStatus, err := handler.Handler(c)
// 	if err != nil {
// 		if data != nil {
// 			commonResponse.FailWithData(err.Error(), data)
// 		} else {
// 			commonResponse.Fail(err.Error())
// 		}
// 		return c.JSON(httpStatus, commonResponse)
// 	}
// 	commonResponse.Success(data)
// 	return c.JSON(httpStatus, commonResponse)
// }
// }
