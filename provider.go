package ethrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
)

type RpcProvider struct {
	Url   string
	index int64
}

func NewRpcProvider(url string) *RpcProvider {
	return &RpcProvider{Url: url}
}

func (p *RpcProvider) requestBody(action string, params []interface{}) *bytes.Buffer {
	index := atomic.AddInt64(&p.index, 1)
	return bytes.NewBuffer(NewRequest(action, params, index).ToBytes())
}

func (p *RpcProvider) sendRequest(body io.Reader) (resp *Response, err error) {
	response, err := http.Post(p.Url, "application/json", body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response http status %s", response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(&resp)
	return
}

type InfuraProvider struct {
	ID     string
	Secret string

	RpcProvider
}

func NewInfuraProvider(id, secret string) *InfuraProvider {
	return &InfuraProvider{id, secret,
		RpcProvider{
			Url: fmt.Sprintf("https://mainnet.infura.io/v3/%s", id),
		},
	}
}
