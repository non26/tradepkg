package bnclient

import (
	"net/http"
)

type IBinanceSerivceHttpClient interface {
	Do(bnreq *http.Request) (*http.Response, error)
	SetClient(t *http.Transport)
}

type binanceSerivceHttpClient struct {
	client *http.Client
}

func NewBinanceSerivceHttpClient() IBinanceSerivceHttpClient {
	binanceClient := binanceSerivceHttpClient{}
	return &binanceClient
}

func (b *binanceSerivceHttpClient) SetClient(t *http.Transport) {
	b.client = &http.Client{
		Transport: t,
	}
}

func (b *binanceSerivceHttpClient) Do(bnreq *http.Request) (*http.Response, error) {
	res, err := b.client.Do(bnreq)
	if err != nil {
		return nil, err
	}
	return res, nil
}
