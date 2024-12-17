package binancecaller

import (
	"fmt"
	"net/http"

	binanceclient "github.com/non26/tradepkg/pkg/bn/binance_client"
	binancemodel "github.com/non26/tradepkg/pkg/bn/binance_model"
	binancerequest "github.com/non26/tradepkg/pkg/bn/binance_request"
	binanceresponse "github.com/non26/tradepkg/pkg/bn/binance_response"
	binanetransport "github.com/non26/tradepkg/pkg/bn/binance_transport"
)

/*
generic Q for requset model
generic P for response model
*/
type ICallBinance[Q, P any] interface {
	NeedSignature(need_signature bool) ICallBinance[Q, P]
	CallBinance(
		request binancemodel.IBnFutureServiceRequest,
		base_url string,
		end_point string,
		method string,
		secret_key string,
		api_key string,
		service_name string,
	) (*P, error)
}

type callBinance[Q any, P any] struct {
	http_request   binancerequest.IBinanceServiceHttpRequest[Q]
	http_response  binanceresponse.IBinanceServiceHttpResponse[P]
	http_transport binanetransport.IBinanceServiceHttpTransport
	http_client    binanceclient.IBinanceSerivceHttpClient
	need_signature bool
}

func NewCallBinance[Q, P any](
	http_request binancerequest.IBinanceServiceHttpRequest[Q],
	http_response binanceresponse.IBinanceServiceHttpResponse[P],
	http_transport binanetransport.IBinanceServiceHttpTransport,
	http_client binanceclient.IBinanceSerivceHttpClient,
) ICallBinance[Q, P] {
	c := callBinance[Q, P]{
		http_request,
		http_response,
		http_transport,
		http_client,
		true,
	}
	return &c
}

func (c *callBinance[Q, P]) NeedSignature(need_signature bool) ICallBinance[Q, P] {
	c.need_signature = need_signature
	return c
}

func (c *callBinance[Q, P]) CallBinance(
	request binancemodel.IBnFutureServiceRequest,
	base_url string,
	end_point string,
	method string,
	secret_key string,
	api_key string,
	service_name string,
) (*P, error) {

	err := c.http_request.NewBinanceHttpRequest(fmt.Sprintf("%v%v", base_url, end_point))
	if err != nil {
		return nil, err
	}

	request.PrepareRequest()
	data := request.GetData().(*Q)
	signature := ""
	if c.need_signature {
		signature, err = c.http_request.CreateRequestSignUrl(data, secret_key)
		if err != nil {
			return nil, err
		}
		c.http_request.AddHeader(api_key)

		switch method {
		case http.MethodPost:
			c.http_request.RequestPostMethod(signature)
		default:
			c.http_request.RequestGetMethod(signature)
		}
	} else {
		c.http_request.GetBinanceRequest().Method = http.MethodGet
	}

	c.http_client.SetClient(c.http_transport.GetTransport())
	err = c.http_client.Do(c.http_request.GetBinanceRequest())
	if err != nil {
		return nil, err
	}

	c.http_response.SetResponse(c.http_client.GetBinanceHttpClientResponse())
	err = c.http_response.DecodeBinanceServiceResponse(service_name)
	if err != nil {
		return nil, err
	}

	return c.http_response.GetBinanceServiceResponse(), nil
}
