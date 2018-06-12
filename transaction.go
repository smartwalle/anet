package anet4go

func (this *AuthorizeNet) CreateTransaction(param *CreateTransactionParam) (result *CreateTransactionRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}