package anet

type TransactionType string

const (
	TransactionTypeCharge              TransactionType = "authCaptureTransaction"
	TransactionTypeAuthOnly            TransactionType = "authOnlyTransaction"
	TransactionTypeCaptureWithTransId  TransactionType = "priorAuthCaptureTransaction"
	TransactionTypeCaptureWithAuthCode TransactionType = "captureOnlyTransaction"
	TransactionTypeRefund              TransactionType = "refundTransaction"
)

type ECheckType string

const (
	ECheckTypePPD ECheckType = "PPD"
	ECheckTypeWEB ECheckType = "WEB"
	ECheckTypeCCD ECheckType = "CCD"
	ECheckTypeTEL ECheckType = "TEL"
	ECheckTypeARC ECheckType = "ARC"
	ECheckTypeBOC ECheckType = "BOC"
)

type AccountType string

const (
	AccountTypeChecking         AccountType = "checking"
	AccountTypeSaving           AccountType = "savings"
	AccountTypeBusinessChecking AccountType = "businessChecking"
)

type CreateTransactionParam struct {
	CreateTransactionRequest struct {
		BasicParam
		TransactionRequest struct {
			TransactionType TransactionType `json:"transactionType,omitempty"`
			Amount          string          `json:"amount,omitempty"`
			Payment         *Payment        `json:"payment,omitempty"`
			LineItems       *LineItems      `json:"lineItems,omitempty"`
			Tax             *Amount         `json:"tax,omitempty"`
			Duty            *Amount         `json:"duty,omitempty"`
			Shipping        *Amount         `json:"shipping,omitempty"`
			PoNumber        string          `json:"poNumber,omitempty"`
			Customer        *Customer       `json:"customer,omitempty"`
			BillTo          *Address        `json:"billTo,omitempty"`
			ShipTo          *Address        `json:"shipTo,omitempty"`
			CustomerIP      string          `json:"customerIP,omitempty"`
			RefTransId      string          `json:"refTransId,omitempty"`
			AuthCode        string          `json:"authCode,omitempty"`
		} `json:"transactionRequest"`
	} `json:"createTransactionRequest"`
}

func (this *CreateTransactionParam) SetMerchantAuthentication(m MerchantAuthentication) {
	this.CreateTransactionRequest.MerchantAuthentication = m
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
		Errors         []*Error   `json:"errors"`
	} `json:"transactionResponse"`
	RefId    string    `json:"refId"`
	Messages *Messages `json:"messages"`
}

type CreditCard struct {
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	CardCode       string `json:"cardCode,omitempty"`
}

type BankAccount struct {
	AccountType   AccountType `json:"accountType,omitempty"`
	RoutingNumber string      `json:"routingNumber,omitempty"`
	AccountNumber string      `json:"accountNumber,omitempty"`
	NameOnAccount string      `json:"nameOnAccount,omitempty"`
	ECheckType    ECheckType  `json:"echeckType,omitempty"`
	BankName      string      `json:"bankName,omitempty"`
	CheckNumber   string      `json:"checkNumber,omitempty"`
}

type Payment struct {
	CreditCard  *CreditCard  `json:"creditCard,omitempty"`
	BankAccount *BankAccount `json:"bankAccount,omitempty"`
}

type Amount struct {
	Amount      string `json:"amount"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Address struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Company     string `json:"company"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	FaxNumber   string `json:"faxNumber,omitempty"`
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
	Type  string `json:"type"`
	Id    string `json:"id"`
	Email string `json:"email"`
}

type TransactionDetailsParam struct {
	GetTransactionDetailsRequest struct {
		BasicParam
		TransId string `json:"transId,omitempty"`
	} `json:"getTransactionDetailsRequest"`
}

func (this *TransactionDetailsParam) SetMerchantAuthentication(m MerchantAuthentication) {
	this.GetTransactionDetailsRequest.MerchantAuthentication = m
}

type TransactionDetailsRsp struct {
	TransactionDetailsResponse struct {
		Messages    *Messages    `json:"messages"`
		Transaction *Transaction `json:"transaction"`
	} `json:"getTransactionDetailsResponse"`
	RefId    string    `json:"refId"`
	Messages *Messages `json:"messages"`
}

type Transaction struct {
	TransId                   string        `json:"transId"`
	RefTransId                string        `json:"refTransId"`
	SplitTenderId             string        `json:"splitTenderId"`
	SubmitTimeUTC             string        `json:"submitTimeUTC"`
	SubmitTimeLocal           string        `json:"submitTimeLocal"`
	TransactionType           string        `json:"transactionType"`
	TransactionStatus         string        `json:"transactionStatus"`
	ResponseCode              string        `json:"responseCode"`
	ResponseReasonCode        string        `json:"responseReasonCode"`
	ResponseReasonDescription string        `json:"responseReasonDescription"`
	AuthCode                  string        `json:"authCode"`
	AVSResponse               string        `json:"AVSResponse"`
	CardCodeResponse          string        `json:"cardCodeResponse"`
	CAVVResponse              string        `json:"CAVVResponse"`
	FDSFilterAction           string        `json:"FDSFilterAction"`
	FDSFilters                *FDSFilters   `json:"FDSFilters"`
	Batch                     *Batch        `json:"batch"`
	Order                     *Order        `json:"order"`
	RequestedAmount           string        `json:"requestedAmount"`
	AuthAmount                string        `json:"authAmount"`
	SettleAmount              string        `json:"settleAmount"`
	Tax                       *Amount       `json:"tax"`
	Shipping                  *Amount       `json:"shipping"`
	Duty                      *Amount       `json:"duty"`
	LineItems                 *LineItems    `json:"lineItems"`
	PrepaidBalanceRemaining   string        `json:"prepaidBalanceRemaining"`
	TaxExempt                 string        `json:"taxExempt"`
	Payment                   *Payment      `json:"payment"`
	Customer                  *Customer     `json:"customer"`
	BillTo                    *Address      `json:"billTo"`
	ShipTo                    *Address      `json:"shipTo"`
	RecurringBilling          string        `json:"recurringBilling"`
	CustomerIP                string        `json:"customerIP"`
	Subscription              *Subscription `json:"subscription"`
	Profile                   *Profile      `json:"Profile"`
}

type FDSFilter struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}

type FDSFilters struct {
	FDSFilter []*FDSFilter `json:"FDSFilter"`
}

type Batch struct {
	BatchId             string `json:"batchId"`
	SettlementTimeUTC   string `json:"settlementTimeUTC"`
	SettlementTimeLocal string `json:"settlementTimeLocal"`
	SettlementState     string `json:"settlementState"`
}

type Order struct {
	InvoiceNumber       string `json:"invoiceNumber"`
	Description         string `json:"description"`
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
}

type Subscription struct {
	Id             string         `json:"id"`
	PayNum         string         `json:"payNum"`
	MarketType     string         `json:"marketType"`
	Product        string         `json:"product"`
	ReturnedItems  *ReturnedItems `json:"returnedItems"`
	Solution       *Solution      `json:"solution"`
	MobileDeviceId string         `json:"mobileDeviceId"`
}

type ReturnedItems struct {
	ReturnedItem *ReturnedItem `json:"returnedItem"`
}

type ReturnedItem struct {
	Id          string `json:"id"`
	DateUTC     string `json:"dateUTC"`
	DateLocal   string `json:"dateLocal"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Solution struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	VendorName string `json:"vendorName"`
}

type Profile struct {
	CustomerProfileId        string `json:"customerProfileId"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId"`
}
