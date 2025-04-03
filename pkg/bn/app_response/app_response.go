package appresponse

var SuccessCode string = "0000"
var FailCode string = "9999"

type CommonResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewCommonResponse(data interface{}) *CommonResponse {
	return &CommonResponse{
		Data: data,
	}
}

func (c *CommonResponse) Fail(message string) {
	c.Code = FailCode
	c.Message = message
}

func (c *CommonResponse) Success() {
	c.Code = SuccessCode
	c.Message = "Success"
}
