package anet_test

import (
	"github.com/smartwalle/anet"
	"os"
	"testing"
)

const (
	kAPILoginId     = "8bSv25KY"
	kTransactionKey = "3D7zFXrt6ebD535V"
)

var client *anet.Client

func TestMain(m *testing.M) {
	client = anet.New(kAPILoginId, kTransactionKey, false)
	os.Exit(m.Run())
}
