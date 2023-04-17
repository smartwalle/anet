package anet

import "fmt"

type Message struct {
	Code        string `json:"code"`
	Text        string `json:"text"`
	Description string `json:"description"`
}

type Messages struct {
	ResultCode string     `json:"resultCode"`
	Message    []*Message `json:"message"`
}

func (this *Messages) Error() string {
	var msg *Message
	if len(this.Message) > 0 {
		msg = this.Message[0]
	}
	if msg != nil {
		return fmt.Sprintf("[%s - %s] %s", this.ResultCode, msg.Code, msg.Text)
	}
	return this.ResultCode
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

type Error struct {
	ErrorCode string `json:"errorCode"`
	ErrorText string `json:"errorText"`
}

func (this *Error) Error() string {
	return fmt.Sprintf("%s - %s", this.ErrorCode, this.ErrorText)
}
