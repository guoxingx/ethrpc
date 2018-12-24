package ethrpc

import (
	"encoding/json"
)

type Request struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      int64       `json:"id"`
	Params  interface{} `json:"params"`
}

func NewRequest(action string, params interface{}, id int64) *Request {
	return &Request{
		JsonRpc: "2.0",
		Method:  action,
		ID:      id,
		Params:  params,
	}
}

func (r *Request) ToBytes() []byte {
	data, _ := json.Marshal(r)
	return data
}
