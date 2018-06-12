package anet4go

const (
	K_TRANSACTION_TYPE_CHARGE                 = "authCaptureTransaction"
	K_TRANSACTION_TYPE_AUTH_ONLY              = "authOnlyTransaction"
	K_TRANSACTION_TYPE_CAPTURE_WITH_TRANS_ID  = "priorAuthCaptureTransaction"
	K_TRANSACTION_TYPE_CAPTURE_WITH_AUTH_CODE = "captureOnlyTransaction"
	K_TRANSACTION_TYPE_REFUND                 = "refundTransaction"
)

type CreateTransactionParam struct {
	CreateTransactionRequest struct {
		BasicParam
		TransactionRequest struct {
			TransactionType string     `json:"transactionType,omitempty"`
			Amount          string     `json:"amount"`
			Payment         *Payment   `json:"payment,omitempty"`
			LineItems       *LineItems `json:"lineItems,omitempty"`
			Tax             *Amount    `json:"tax,omitempty"`
			Duty            *Amount    `json:"duty,omitempty"`
			Shipping        *Amount    `json:"shipping,omitempty"`
			PoNumber        string     `json:"poNumber,omitempty"`
			Customer        *Customer  `json:"customer,omitempty"`
			BillTo          *Address   `json:"billTo,omitempty"`
			ShipTo          *Address   `json:"shipTo,omitempty"`
			CustomerIP      string     `json:"customerIP,omitempty"`
			RefTransId      string     `json:"refTransId,omitempty"`
			AuthCode        string     `json:"authCode,omitempty"`
		} `json:"transactionRequest"`
	} `json:"createTransactionRequest"`
}

func (this *CreateTransactionParam) SetMerchantAuthentication(m MerchantAuthentication) {
	this.CreateTransactionRequest.MerchantAuthentication = m
}

type CreditCard struct {
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	CardCode       string `json:"cardCode,omitempty"`
}

type Payment struct {
	CreditCard *CreditCard `json:"creditCard,omitempty"`
}

type Amount struct {
	Amount      string `json:"amount"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Address struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Company   string `json:"company"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

type LineItem struct {
	ItemId      string `json:"itemId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	UnitPrice   string `json:"unitPrice"`
}

type LineItems struct {
	LineItem *LineItem `json:"lineItem,omitempty"`
}

type Customer struct {
	Id string `json:"id"`
}

type TransactionRsp struct {
	TransactionResponse struct {
		ResponseCode   string     `json:"responseCode"`
		AuthCode       string     `json:"authCode"`
		AVSResultCode  string     `json:"avsResultCode"`
		CVVResultCode  string     `json:"cvvResultCode"`
		CAVVResultCode string     `json:"cavvResultCode"`
		TransId        string     `json:"transId"`
		RefTransID     string     `json:"refTransID"`
		TransHash      string     `json:"transHash"`
		AccountNumber  string     `json:"accountNumber"`
		AccountType    string     `json:"accountType"`
		Messages       []*Message `json:"messages"`
	} `json:"transactionResponse"`
	RefId    string    `json:"refId"`
	Messages *Messages `json:"messages"`
}
