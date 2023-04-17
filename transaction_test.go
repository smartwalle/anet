package anet_test

import (
	"fmt"
	"github.com/smartwalle/anet"
	"testing"
)

func TestAuthorizeNet_CreateTransaction(t *testing.T) {
	var p = &anet.CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = anet.TransactionTypeCharge
	p.CreateTransactionRequest.TransactionRequest.Amount = "100"
	var payment = &anet.Payment{}
	var creditCard = &anet.CreditCard{}
	creditCard.CardNumber = "5145919892544954"
	creditCard.ExpirationDate = "2020-01"
	creditCard.CardCode = "123"
	payment.CreditCard = creditCard
	p.CreateTransactionRequest.TransactionRequest.PoNumber = "123456789011"
	p.CreateTransactionRequest.TransactionRequest.Payment = payment

	var billTo = &anet.Address{}
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
	fmt.Println(rsp.TransactionResponse.TransId)
}

func TestAuthorizeNet_Charge(t *testing.T) {
	//fmt.Println(client.ChargeWithCreditCard("12342", "100", "5145919892544954", "2020-01", "123", "f", "y", "China", "SC", "CD", "123456", "123456", "123456", "ADD"))
}

func TestAuthorizeNet_Auth(t *testing.T) {
	//fmt.Println(client.AuthWithCreditCard("12343", "100", "5145919892544954", "2020-01", "123", "f", "y", "China", "SC", "CD", "123456", "123456", "123456", "ADD"))
}

func TestAuthorizeNet_CaptureWithTransId(t *testing.T) {
	//fmt.Println(client.CaptureWithTransId("60106662170", "300"))
}

func TestAuthorizeNet_CaptureWithAutoCode(t *testing.T) {
	//	client.CaptureWithAutoCode("MFQFPQ", "100", "5145919892544954", "2020-01", "123")
}

func TestAuthorizeNet_Refund(t *testing.T) {
	//fmt.Println(client.RefundWithCreditCard("60106662016", "100", "5145919892544954", "2020-01"))
}

func TestAuthorizeNet_GetTransactionDetails(t *testing.T) {
	//	fmt.Println(client.GetTransactionDetails("60104455428"))
}

func TestAuthorizeNet_DebitWithBankAccount(t *testing.T) {
	//fmt.Println(client.DebitWithBankAccount("1234", "1000", AccountTypeChecking, "121042882", "123456789", "John Doe", "", "", "", "", "", "", "", "", "", "", "", ""))
}
