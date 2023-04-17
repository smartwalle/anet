package anet

type AuthenticateTestParam struct {
	AuthenticateTestRequest struct {
		BasicParam
	} `json:"authenticateTestRequest"`
}

func (this *AuthenticateTestParam) SetMerchantAuthentication(m MerchantAuthentication) {
	this.AuthenticateTestRequest.MerchantAuthentication = m
}

type AuthenticateTestRsp struct {
	RefId    string    `json:"refId"`
	Messages *Messages `json:"messages"`
}
