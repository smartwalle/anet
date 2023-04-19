package anet

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	kProductionURL = "https://api.authorize.net/xml/v1/request.api"
	kSandboxURL    = "https://apitest.authorize.net/xml/v1/request.api"
)

var (
	kResponsePrefix = []byte("\xef\xbb\xbf")
)

type Client struct {
	apiDomain      string
	apiLoginId     string
	transactionKey string
	authentication MerchantAuthentication
	Client         *http.Client
}

func New(apiLoginId, transactionKey string, isProduction bool) (client *Client) {
	client = &Client{}
	client.Client = http.DefaultClient
	client.apiLoginId = apiLoginId
	client.transactionKey = transactionKey
	client.authentication = MerchantAuthentication{Name: apiLoginId, TransactionKey: transactionKey}
	if isProduction {
		client.apiDomain = kProductionURL
	} else {
		client.apiDomain = kSandboxURL
	}
	return client
}

func (this *Client) doRequest(method string, param Param, results interface{}) error {
	var body io.Reader
	if param != nil {
		param.SetMerchantAuthentication(this.authentication)
		pData, err := json.Marshal(param)
		if err != nil {
			return err
		}
		body = bytes.NewReader(pData)
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
