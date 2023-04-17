package anet

func (this *Client) CreateTransaction(param *CreateTransactionParam) (result *TransactionRsp, err error) {
	err = this.doRequest("POST", param, &result)
	if result.Messages.ResultCode != "Ok" && err == nil {
		if len(result.TransactionResponse.Errors) > 0 {
			err = result.TransactionResponse.Errors[0]
		} else {
			err = result.Messages
		}
	}
	return result, err
}

func (this *Client) ChargeWithCreditCard(poNumber, amount, cardNumber, expirationDate, cardCode, firstName, lastName, country, state, city, zip, phone, fax, address string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.PoNumber = poNumber
	p.CreateTransactionRequest.TransactionRequest.TransactionType = TransactionTypeCharge
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
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
	billTo.PhoneNumber = phone
	billTo.FaxNumber = fax
	p.CreateTransactionRequest.TransactionRequest.BillTo = billTo
	return this.CreateTransaction(p)
}

func (this *Client) AuthWithCreditCard(poNumber, amount, cardNumber, expirationDate, cardCode, firstName, lastName, country, state, city, zip, phone, fax, address string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.PoNumber = poNumber
	p.CreateTransactionRequest.TransactionRequest.TransactionType = TransactionTypeAuthOnly
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
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
	billTo.PhoneNumber = phone
	billTo.FaxNumber = fax
	p.CreateTransactionRequest.TransactionRequest.BillTo = billTo
	return this.CreateTransaction(p)
}

func (this *Client) CaptureWithTransId(transId, amount string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = TransactionTypeCaptureWithTransId
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
	p.CreateTransactionRequest.TransactionRequest.RefTransId = transId
	return this.CreateTransaction(p)
}

func (this *Client) CaptureWithAutoCode(authCode, amount, cardNumber, expirationDate, cardCode string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = TransactionTypeCaptureWithAuthCode
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

func (this *Client) RefundWithCreditCard(transId, amount, cardNumber, expirationDate string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = TransactionTypeRefund
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

func (this *Client) DebitWithBankAccount(poNumber, amount string, accountType AccountType, routingNumber, accountNumber, nameOnAccount string, eCheckType ECheckType, bankName, checkNumber, firstName, lastName, country, state, city, zip, phone, fax, address string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.PoNumber = poNumber
	p.CreateTransactionRequest.TransactionRequest.TransactionType = TransactionTypeCharge
	p.CreateTransactionRequest.TransactionRequest.Amount = amount
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
	billTo.PhoneNumber = phone
	billTo.FaxNumber = fax
	p.CreateTransactionRequest.TransactionRequest.BillTo = billTo
	return this.CreateTransaction(p)
}

func (this *Client) CreditWithBankAccount(transId, amount string, accountType AccountType, routingNumber, accountNumber, nameOnAccount string) (result *TransactionRsp, err error) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = TransactionTypeRefund
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

func (this *Client) GetTransactionDetails(transId string) (result *TransactionDetailsRsp, err error) {
	var param = &TransactionDetailsParam{}
	param.GetTransactionDetailsRequest.TransId = transId
	err = this.doRequest("POST", param, &result)
	if result.Messages.ResultCode != "Ok" && err == nil {
		err = result.Messages
	}
	return result, err
}
