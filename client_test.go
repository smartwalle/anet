package anet_test

import (
	"github.com/smartwalle/anet"
	"os"
	"testing"
)

const (
	kAPILoginId     = "4Wb9uKQ8L"
	kTransactionKey = "43H7ncQDc65979ky"
)

var client *anet.Client

func TestMain(m *testing.M) {
	client = anet.New(kAPILoginId, kTransactionKey, false)
	os.Exit(m.Run())
}
