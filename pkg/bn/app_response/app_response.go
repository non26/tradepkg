package appresponse

var SuccessCode string = "0000"
var FailCode string = "9999"
var FailMsg string = "Failed"
var SuccessMsg string = "Success"
var Unknown string = "unknown"

type CommonResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewCommonResponse() *CommonResponse {
	return &CommonResponse{}
}

func (c *CommonResponse) Fail(message string) {
	c.Code = FailCode
	if message == "" {
		c.Message = FailMsg
	} else {
		c.Message = message
	}
}

func (c *CommonResponse) FailWithData(message string, data interface{}) {
	c.Fail(message)
	c.Data = data
}

func (c *CommonResponse) Success(data interface{}) {
	c.Code = SuccessCode
	c.Message = SuccessMsg
	c.Data = data
}
