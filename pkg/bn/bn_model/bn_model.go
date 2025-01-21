package bnmodel

type IBnFutureServiceRequest interface {
	PrepareRequest()
	GetData() interface{}
}
type ITest interface {
	IBnFutureServiceRequest
	Test()
}
