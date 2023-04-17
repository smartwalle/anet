package anet_test

import (
	"testing"
)

func TestAuthorize_AuthenticateTest(t *testing.T) {
	var rsp, err = client.AuthenticateTest()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp.Messages.ResultCode)
}
