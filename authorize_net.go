package anet4go

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
)

// 测试信息

// API Credentials
// API LOGIN ID
// 4Wb9uKQ8L
//
// TRANSACTION KEY
// 43H7ncQDc65979ky
//
// KEY
// Simon

const (
	k_PRODUCTION_API_URL = "https://api.authorize.net/xml/v1/request.api"
	k_SANDBOX_API_URL    = "https://apitest.authorize.net/xml/v1/request.api"
)

var (
	k_RESPONSE_PREFIX = []byte("\xef\xbb\xbf")
)

type AuthorizeNet struct {
	apiDomain      string
	apiLoginId     string
	transactionKey string
	m              MerchantAuthentication
	Client         *http.Client
}

func New(apiLoginId, transactionKey string, isProduction bool) (client *AuthorizeNet) {
	client = &AuthorizeNet{}
	client.Client = http.DefaultClient
	client.apiLoginId = apiLoginId
	client.transactionKey = transactionKey
	client.m = MerchantAuthentication{Name: apiLoginId, TransactionKey: transactionKey}
	if isProduction {
		client.apiDomain = k_PRODUCTION_API_URL
	} else {
		client.apiDomain = k_SANDBOX_API_URL
	}
	return client
}

func (this *AuthorizeNet) doRequest(method string, param Param, results interface{}) (err error) {
	var buf io.Reader
	if param != nil {
		param.SetMerchantAuthentication(this.m)
		p, err := json.Marshal(param)
		if err != nil {
			return err
		}
		buf = strings.NewReader(string(p))
	}

	req, err := http.NewRequest(method, this.apiDomain, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := this.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	data = bytes.TrimPrefix(data, k_RESPONSE_PREFIX)

	fmt.Println(string(data))

	err = json.Unmarshal(data, results)
	if err != nil {
		return err
	}

	return err
}
