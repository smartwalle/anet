package anet_test

import (
	"github.com/smartwalle/anet"
	"testing"
)

func TestClient_CreateTransaction(t *testing.T) {
	var p = &anet.CreateTransactionParam{}
	p.CreateTransactionRequest.TransactionRequest.TransactionType = anet.TransactionTypeCharge
	p.CreateTransactionRequest.TransactionRequest.Amount = "100"
	var payment = &anet.Payment{}
	var creditCard = &anet.CreditCard{}
	creditCard.CardNumber = "5145919892544954"
	creditCard.ExpirationDate = "2024-01"
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
	t.Log(rsp.TransactionResponse.TransId)
}

func TestClient_Charge(t *testing.T) {
	t.Log(client.ChargeWithCreditCard("12342", "100", "5145919892544954", "2024-01", "123", "f", "y", "China", "SC", "CD", "123456", "123456", "123456", "ADD"))
}

func TestClient_Auth(t *testing.T) {
	t.Log(client.AuthWithCreditCard("12343", "100", "5145919892544954", "2024-01", "123", "f", "y", "China", "SC", "CD", "123456", "123456", "123456", "ADD"))
}

func TestClient_CaptureWithTransId(t *testing.T) {
	t.Log(client.CaptureWithTransId("60106662170", "300"))
}

func TestClient_CaptureWithAuthCode(t *testing.T) {
	client.CaptureWithAuthCode("MFQFPQ", "100", "5145919892544954", "2024-01", "123")
}

func TestClient_Refund(t *testing.T) {
	t.Log(client.RefundWithCreditCard("60106662016", "100", "5145919892544954", "2024-01"))
}

func TestClient_GetTransactionDetails(t *testing.T) {
	t.Log(client.GetTransactionDetails("60104455428"))
}

func TestClient_DebitWithBankAccount(t *testing.T) {
	t.Log(client.DebitWithBankAccount("1234", "1000", anet.AccountTypeChecking, "121042882", "123456789", "John Doe", "", "", "", "", "", "", "", "", "", "", "", ""))
}
