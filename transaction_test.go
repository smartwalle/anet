package anet4go

import (
	"testing"
	"fmt"
)

func TestAuthorizeNet_CreateTransaction(t *testing.T) {
	var p = &CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = K_TRANSACTION_TYPE_AUTH_ONLY
	p.CreateTransactionRequest.TransactionRequest.Amount = "100"
	//p.CreateTransactionRequest.TransactionRequest.RefTransId = "60104453508"
	var payment = &Payment{}
	var creditCard = &CreditCard{}
	creditCard.CardNumber = "5145919892544954"
	creditCard.ExpirationDate = "2020-01"
	creditCard.CardCode = "123"
	payment.CreditCard = creditCard
	p.CreateTransactionRequest.TransactionRequest.PoNumber = "123456789"
	p.CreateTransactionRequest.TransactionRequest.Payment = payment

	var billTo = &Address{}
	billTo.Address = "ADD"
	billTo.FirstName = "Feng"
	billTo.LastName = "Yang"
	billTo.Country = "China"
	billTo.City = "CD"
	billTo.State = "SC"
	billTo.Zip = "123456"
	p.CreateTransactionRequest.TransactionRequest.BillTo = billTo

	var rsp, err = client.CreateTransaction(p)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rsp.Messages.ResultCode)
}