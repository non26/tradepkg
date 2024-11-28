package binancemodel

type IBnFutureServiceRequest interface {
	PrepareRequest()
	GetData() interface{}
}
type ITest interface {
	IBnFutureServiceRequest
	Test()
}
