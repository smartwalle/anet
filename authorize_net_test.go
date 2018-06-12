package anet4go

import (
	"os"
	"testing"
)

const (
	k_API_LOGIN_ID    = "4Wb9uKQ8L"
	k_TRANSACTION_KEY = "43H7ncQDc65979ky"
)

var client *AuthorizeNet

func TestMain(m *testing.M) {
	client = New(k_API_LOGIN_ID, k_TRANSACTION_KEY, false)
	os.Exit(m.Run())
}
