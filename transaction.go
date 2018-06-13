package anet4go

func (this *AuthorizeNet) CreateTransaction(param *CreateTransactionParam) (result *TransactionRsp, err error) {
	err = this.doRequest("POST", param, &result)
	if result.Messages.ResultCode != "Ok" && err == nil {
		err = result.Messages
	}
	return result, err
}

func (this *AuthorizeNet) ChargeWithCreditCard(poNumber, amount, firstName, lastName, cardNumber, expirationDate, cardCode, country, state, city, zip, address string) (result *TransactionRsp, err error) {
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

func (this *AuthorizeNet) AuthWithCreditCard(poNumber, amount, firstName, lastName, cardNumber, expirationDate, cardCode, country, state, city, zip, address string) (result *TransactionRsp, err error) {
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
	p.CreateTransactionRequest.TransactionRequest.RefTransId = transId
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

func (this *AuthorizeNet) RefundWithCreditCard(transId, amount, cardNumber, expirationDate string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = K_TRANSACTION_TYPE_REFUND
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
	var payment = &Payment{}
	var creditCard = &CreditCard{}
	creditCard.CardNumber = cardNumber
	creditCard.ExpirationDate = expirationDate
	payment.CreditCard = creditCard
	p.CreateTransactionRequest.TransactionRequest.Payment = payment
	p.CreateTransactionRequest.TransactionRequest.RefTransId = transId
	return this.CreateTransaction(p)
}

func (this *AuthorizeNet) DebitWithBankAccount(poNumber, amount, accountType, routingNumber, accountNumber, nameOnAccount, eCheckType, bankName, checkNumber, firstName, lastName, country, state, city, zip, address string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.PoNumber = poNumber
	p.CreateTransactionRequest.TransactionRequest.TransactionType = K_TRANSACTION_TYPE_CHARGE
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
	p.CreateTransactionRequest.RefId = poNumber
	var payment = &Payment{}
	var bankAccount = &BankAccount{}
	bankAccount.AccountType = accountType
	bankAccount.RoutingNumber = routingNumber
	bankAccount.AccountNumber = accountNumber
	bankAccount.NameOnAccount = nameOnAccount
	bankAccount.ECheckType = eCheckType
	bankAccount.BankName = bankName
	bankAccount.CheckNumber = checkNumber
	payment.BankAccount = bankAccount
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

func (this *AuthorizeNet) CreditWithBankAccount(transId, amount, accountType, routingNumber, accountNumber, nameOnAccount string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = K_TRANSACTION_TYPE_REFUND
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
	var payment = &Payment{}
	var bankAccount = &BankAccount{}
	bankAccount.AccountType = accountType
	bankAccount.RoutingNumber = routingNumber
	bankAccount.AccountNumber = accountNumber
	bankAccount.NameOnAccount = nameOnAccount
	payment.BankAccount = bankAccount
	p.CreateTransactionRequest.TransactionRequest.Payment = payment
	p.CreateTransactionRequest.TransactionRequest.RefTransId = transId
	return this.CreateTransaction(p)
}

func (this *AuthorizeNet) GetTransactionDetails(transId string) (result *TransactionDetailsRsp, err error) {
	var param = &TransactionDetailsParam{}
	param.GetTransactionDetailsRequest.TransId = transId
	err = this.doRequest("POST", param, &result)
	if result.Messages.ResultCode != "Ok" && err == nil {
		err = result.Messages
	}
	return result, err
}
