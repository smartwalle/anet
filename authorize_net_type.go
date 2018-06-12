package anet4go

type Message struct {
	Code        string `json:"code"`
	Text        string `json:"text"`
	Description string `json:"description"`
}

type Messages struct {
	ResultCode string     `json:"resultCode"`
	Message    []*Message `json:"message"`
}

type MerchantAuthentication struct {
	Name           string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}

type Param interface {
	SetMerchantAuthentication(m MerchantAuthentication)
}

type BasicParam struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	RefId                  string                 `json:"refId"`
}
