package infura

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
)

// infura.io with http request
// Ethereum only
type InfuraProvider struct {
	ID     string
	Secret string

	index int64
}

func (p *InfuraProvider) requestBody(action string, params []interface{}) *bytes.Buffer {
	index := atomic.AddInt64(&p.index, 1)
	return bytes.NewBuffer(NewRequest(action, params, index).ToBytes())
}

func (p *InfuraProvider) sendRequest(body io.Reader) (resp *Response, err error) {
	url := fmt.Sprintf("https://mainnet.infura.io/v3/%s", p.ID)
	response, err := http.Post(url, "application/json", body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response http status %s", response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(&resp)
	return
}
