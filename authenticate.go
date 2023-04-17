package anet

func (this *Client) AuthenticateTest() (result *AuthenticateTestRsp, err error) {
	var param = &AuthenticateTestParam{}
	err = this.doRequest("POST", param, &result)
	return result, err
}
