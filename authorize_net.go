package anet

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	kProductionAPIURL = "https://api.authorize.net/xml/v1/request.api"
	kSandboxAPIURL    = "https://apitest.authorize.net/xml/v1/request.api"
)

var (
	kResponsePrefix = []byte("\xef\xbb\xbf")
)

type Client struct {
	apiDomain      string
	apiLoginId     string
	transactionKey string
	m              MerchantAuthentication
	Client         *http.Client
}

func New(apiLoginId, transactionKey string, isProduction bool) (client *Client) {
	client = &Client{}
	client.Client = http.DefaultClient
	client.apiLoginId = apiLoginId
	client.transactionKey = transactionKey
	client.m = MerchantAuthentication{Name: apiLoginId, TransactionKey: transactionKey}
	if isProduction {
		client.apiDomain = kProductionAPIURL
	} else {
		client.apiDomain = kSandboxAPIURL
	}
	return client
}

func (this *Client) doRequest(method string, param Param, results interface{}) (err error) {
	var body io.Reader
	if param != nil {
		param.SetMerchantAuthentication(this.m)
		p, err := json.Marshal(param)
		if err != nil {
			return err
		}
		body = strings.NewReader(string(p))
	}

	req, err := http.NewRequest(method, this.apiDomain, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	rsp, err := this.Client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	data = bytes.TrimPrefix(data, kResponsePrefix)

	if err = json.Unmarshal(data, results); err != nil {
		return err
	}

	return err
}
