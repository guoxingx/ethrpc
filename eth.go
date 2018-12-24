package ethrpc

import (
	"encoding/json"
	"fmt"
)

func (p *RpcProvider) EthAccounts() (accounts []string, err error) {
	body := p.requestBody("eth_accounts", []interface{}{})
	resp, err := p.sendRequest(body)
	if err != nil {
		return accounts, err
	}
	err = json.Unmarshal(resp.Result, &accounts)
	return
}

func (p *RpcProvider) EthBlockNumber() (height int64, err error) {
	body := p.requestBody("eth_blockNumber", []interface{}{})
	resp, err := p.sendRequest(body)
	if err != nil {
		return height, err
	}
	err = json.Unmarshal(resp.Result, &height)
	return
}

func (p *RpcProvider) EthCall() {
	data := []interface{}{
		map[string]string{
			"from":     "",
			"to":       "",
			"gas":      "",
			"gasPrice": "",
			"value":    "",
			"data":     "",
		},
		"latest",
	}
	panic(data)
}
func (p *RpcProvider) EthEstimateGas()    {}
func (p *RpcProvider) EthGasPrice()       {}
func (p *RpcProvider) EthGetBalance()     {}
func (p *RpcProvider) EthGetBlockByHash() {}

func (p *RpcProvider) EthGetBlockByNumber(height interface{}) (block *BlockResponse, err error) {
	var heightHex string

	switch height.(type) {
	case string:
		if height == "latest" {
			heightHex = height.(string)
		} else if height == "earliest" {
			heightHex = height.(string)
		} else if height == "pending" {
			heightHex = height.(string)
		}
	case int64:
		heightHex = fmt.Sprintf("0x%x", height.(int64))
	default:
		return nil, fmt.Errorf("unsported type for height, value: %v", height)
	}
	if heightHex == "" {
		return nil, fmt.Errorf("failed to parse height to hex, value: %v", height)
	}

	body := p.requestBody("eth_getBlockByNumber", []interface{}{heightHex, false})
	resp, err := p.sendRequest(body)
	if err != nil {
		return nil, err
	}

	var blockHex BlockResponseHex
	err = json.Unmarshal(resp.Result, &blockHex)
	return blockHex.UnHexed(), err
}

func (p *RpcProvider) EthGetBlockTransactionCountByHash()      {}
func (p *RpcProvider) EthGetBlockTransactionCountByNumber()    {}
func (p *RpcProvider) EthGetCode()                             {}
func (p *RpcProvider) EthGetLogs()                             {}
func (p *RpcProvider) EthGetStorageAt()                        {}
func (p *RpcProvider) EthGetTransactionByBlockHashAndIndex()   {}
func (p *RpcProvider) EthGetTransactionByBlockNumberAndIndex() {}
func (p *RpcProvider) EthGetTransactionByHash()                {}
func (p *RpcProvider) EthGetTransactionCount()                 {}
func (p *RpcProvider) EthGetTransactionReceipt()               {}
func (p *RpcProvider) EthGetUncleByBlockHashAndIndex()         {}
func (p *RpcProvider) EthGetUncleByBlockNumberAndIndex()       {}
func (p *RpcProvider) EthGetUncleCountByBlockHash()            {}
func (p *RpcProvider) EthGetUncleCountByBlockNumber()          {}
func (p *RpcProvider) EthGetWork()                             {}
func (p *RpcProvider) EthHashrate()                            {}
func (p *RpcProvider) EthMining()                              {}
func (p *RpcProvider) EthProtocolVersion()                     {}
func (p *RpcProvider) EthSendRawTransaction()                  {}
func (p *RpcProvider) EthSubmitWork()                          {}
func (p *RpcProvider) EthSyncing()                             {}
func (p *RpcProvider) EthListening()                           {}
func (p *RpcProvider) EthPeerCount()                           {}
func (p *RpcProvider) EthVersion()                             {}
func (p *RpcProvider) EthClientVersion()                       {}
