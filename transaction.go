package anet4go

func (this *AuthorizeNet) CreateTransaction(param *CreateTransactionParam) (result *TransactionRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

func (this *AuthorizeNet) Charge(poNumber, amount, firstName, lastName, cardNumber, expirationDate, cardCode, country, state, city, zip, address string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.PoNumber = poNumber
	p.CreateTransactionRequest.TransactionRequest.TransactionType = K_TRANSACTION_TYPE_CHARGE
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
	p.CreateTransactionRequest.RefId = poNumber
	var payment = &Payment{}
	var creditCard = &CreditCard{}
	creditCard.CardNumber = cardNumber
	creditCard.ExpirationDate = expirationDate
	creditCard.CardCode = cardCode
	payment.CreditCard = creditCard
	p.CreateTransactionRequest.TransactionRequest.Payment = payment

	var billTo = &Address{}
	billTo.Address = address
	billTo.FirstName = firstName
	billTo.LastName = lastName
	billTo.Country = country
	billTo.City = city
	billTo.State = state
	billTo.Zip = zip
	p.CreateTransactionRequest.TransactionRequest.BillTo = billTo
	return this.CreateTransaction(p)
}

func (this *AuthorizeNet) Auth(poNumber, amount, firstName, lastName, cardNumber, expirationDate, cardCode, country, state, city, zip, address string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.PoNumber = poNumber
	p.CreateTransactionRequest.TransactionRequest.TransactionType = K_TRANSACTION_TYPE_AUTH_ONLY
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
	p.CreateTransactionRequest.RefId = poNumber
	var payment = &Payment{}
	var creditCard = &CreditCard{}
	creditCard.CardNumber = cardNumber
	creditCard.ExpirationDate = expirationDate
	creditCard.CardCode = cardCode
	payment.CreditCard = creditCard
	p.CreateTransactionRequest.TransactionRequest.Payment = payment

	var billTo = &Address{}
	billTo.Address = address
	billTo.FirstName = firstName
	billTo.LastName = lastName
	billTo.Country = country
	billTo.City = city
	billTo.State = state
	billTo.Zip = zip
	p.CreateTransactionRequest.TransactionRequest.BillTo = billTo
	return this.CreateTransaction(p)
}

func (this *AuthorizeNet) CaptureWithTransId(transId, amount string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = K_TRANSACTION_TYPE_CAPTURE_WITH_TRANS_ID
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
	return this.CreateTransaction(p)
}

func (this *AuthorizeNet) CaptureWithAutoCode(authCode, amount, cardNumber, expirationDate, cardCode string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = K_TRANSACTION_TYPE_CAPTURE_WITH_AUTH_CODE
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
	var payment = &Payment{}
	var creditCard = &CreditCard{}
	creditCard.CardNumber = cardNumber
	creditCard.ExpirationDate = expirationDate
	creditCard.CardCode = cardCode
	payment.CreditCard = creditCard
	p.CreateTransactionRequest.TransactionRequest.Payment = payment
	p.CreateTransactionRequest.TransactionRequest.AuthCode = authCode
	return this.CreateTransaction(p)
}
