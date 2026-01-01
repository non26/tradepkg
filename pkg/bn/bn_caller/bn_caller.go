package bncaller

import (
	"fmt"
	"net/http"

	binanceclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	binancerequest "github.com/non26/tradepkg/pkg/bn/bn_request"
	binanceresponse "github.com/non26/tradepkg/pkg/bn/bn_response"
	binanetransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

/*
generic Q for requset model
generic P for response model
*/
type ICallBinance[Q, P any] interface {
	CallBinance(
		request binancerequest.IBnFutureServiceRequest,
		base_url string,
		end_point string,
		method string,
		secret_key string,
		api_key string,
	) (*P, error)
}

type callBinance[Q any, P any] struct {
	http_request   binancerequest.IBinanceServiceHttpRequest[Q]
	http_response  binanceresponse.IBinanceServiceHttpResponse[P]
	http_transport binanetransport.IBinanceServiceHttpTransport
	http_client    binanceclient.IBinanceSerivceHttpClient
	need_signature bool
	future         bool
	spot           bool
}

// q = is the model for Binance request
// p = is the model for Binance response
func NewCallBinance[Q, P any](
	http_request binancerequest.IBinanceServiceHttpRequest[Q],
	http_response binanceresponse.IBinanceServiceHttpResponse[P],
	http_transport binanetransport.IBinanceServiceHttpTransport,
	http_client binanceclient.IBinanceSerivceHttpClient,
	need_signature bool,
	future bool,
	spot bool,
) ICallBinance[Q, P] {
	if future && spot {
		panic("future and spot cannot be true at the same time")
	}
	c := callBinance[Q, P]{
		http_request,
		http_response,
		http_transport,
		http_client,
		need_signature,
		future,
		spot,
	}
	return &c
}

// q = is the model for Binance request
// p = is the model for Binance response
func (c *callBinance[Q, P]) CallBinance(
	request binancerequest.IBnFutureServiceRequest,
	base_url string,
	end_point string,
	method string,
	secret_key string,
	api_key string,
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
		if c.future {
			c.http_request.AddBnDefaultFutureHeader(api_key)
		} else if c.spot {
			c.http_request.AddBnDefaultSpotHeader(api_key)
		}

		switch method {
		case http.MethodPost:
			c.http_request.RequestPostMethod()
		default:
			c.http_request.RequestGetMethod()
		}
		c.http_request.SetQueryString(signature)
	} else {
		queryString := utils.CreateQueryStringFrom(data)
		c.http_request.SetQueryString(queryString)
		c.http_request.RequestGetMethod()
	}

	c.http_client.SetClient(c.http_transport.GetTransport())
	httpResponse, err := c.http_client.Do(c.http_request.GetBinanceRequest())
	if err != nil {
		return nil, err
	}

	c.http_response.SetResponse(httpResponse)
	err = c.http_response.DecodeBinanceServiceResponse()
	if err != nil {
		return nil, err
	}

	return c.http_response.GetBinanceServiceResponse(), nil
}
