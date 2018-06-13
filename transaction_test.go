package anet4go

//func TestAuthorizeNet_CreateTransaction(t *testing.T) {
//	var p = &CreateTransactionParam{}
//	p.CreateTransactionRequest.TransactionRequest.TransactionType = K_TRANSACTION_TYPE_CHARGE
//	p.CreateTransactionRequest.TransactionRequest.Amount = "100"
//	var payment = &Payment{}
//	var creditCard = &CreditCard{}
//	creditCard.CardNumber = "5145919892544954"
//	creditCard.ExpirationDate = "2020-01"
//	creditCard.CardCode = "123"
//	payment.CreditCard = creditCard
//	p.CreateTransactionRequest.TransactionRequest.PoNumber = "123456789011"
//	p.CreateTransactionRequest.TransactionRequest.Payment = payment
//
//	var billTo = &Address{}
//	billTo.Address = "ADD"
//	billTo.FirstName = "Feng"
//	billTo.LastName = "Yang"
//	billTo.Country = "China"
//	billTo.City = "CD"
//	billTo.State = "SC"
//	billTo.Zip = "123456"
//	p.CreateTransactionRequest.TransactionRequest.BillTo = billTo
//
//	var rsp, err = client.CreateTransaction(p)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println(rsp.TransactionResponse.TransId)
//}
//
//func TestAuthorizeNet_Charge(t *testing.T) {
//	fmt.Println(client.ChargeWithCreditCard("1234", "100", "f", "y", "5145919892544954", "2020-01", "123", "China", "SC", "CD", "123456", "ADD"))
//}
//
//func TestAuthorizeNet_Auth(t *testing.T) {
//	fmt.Println(client.AuthWithCreditCard("1234", "100", "f", "y", "5145919892544954", "2020-01", "123", "China", "SC", "CD", "123456", "ADD"))
//}
//
//func TestAuthorizeNet_CaptureWithTransId(t *testing.T) {
//	fmt.Println(client.CaptureWithTransId("60104495440", "100"))
//}
//
//func TestAuthorizeNet_CaptureWithAutoCode(t *testing.T) {
//	client.CaptureWithAutoCode("MFQFPQ", "100", "5145919892544954", "2020-01", "123")
//}
//
//func TestAuthorizeNet_Refund(t *testing.T) {
//	client.RefundWithCreditCard("60104455428", "100", "4954", "2020-01")
//}
//
//func TestAuthorizeNet_GetTransactionDetails(t *testing.T) {
//	fmt.Println(client.GetTransactionDetails("60104455428"))
//}
//
//func TestAuthorizeNet_DebitWithBankAccount(t *testing.T) {
//	fmt.Println(client.DebitWithBankAccount("1234", "1000", K_ACCOUNT_TYPE_CHECKING, "121042882", "123456789", "John Doe", "", "", "", "", "", "", ""))
//}
