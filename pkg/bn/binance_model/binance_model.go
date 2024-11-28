package binancemodel

type IBnFutureServiceRequest interface {
	PrepareRequest()
	GetData() interface{}
}
