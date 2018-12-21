package infura

import (
	"encoding/json"
	"fmt"
)

func (p *InfuraProvider) EthAccounts() {}

func (p *InfuraProvider) EthBlockNumber() (height int64, err error) {
	body := p.requestBody("eth_blockNumber", []interface{}{})
	resp, err := p.sendRequest(body)
	if err != nil {
		return height, err
	}
	err = json.Unmarshal(resp.Result, &height)
	return
}

func (p *InfuraProvider) EthCall()           {}
func (p *InfuraProvider) EthEstimateGas()    {}
func (p *InfuraProvider) EthGasPrice()       {}
func (p *InfuraProvider) EthGetBalance()     {}
func (p *InfuraProvider) EthGetBlockByHash() {}

func (p *InfuraProvider) EthGetBlockByNumber(height interface{}) (block *BlockResponse, err error) {
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

func (p *InfuraProvider) EthGetBlockTransactionCountByHash()      {}
func (p *InfuraProvider) EthGetBlockTransactionCountByNumber()    {}
func (p *InfuraProvider) EthGetCode()                             {}
func (p *InfuraProvider) EthGetLogs()                             {}
func (p *InfuraProvider) EthGetStorageAt()                        {}
func (p *InfuraProvider) EthGetTransactionByBlockHashAndIndex()   {}
func (p *InfuraProvider) EthGetTransactionByBlockNumberAndIndex() {}
func (p *InfuraProvider) EthGetTransactionByHash()                {}
func (p *InfuraProvider) EthGetTransactionCount()                 {}
func (p *InfuraProvider) EthGetTransactionReceipt()               {}
func (p *InfuraProvider) EthGetUncleByBlockHashAndIndex()         {}
func (p *InfuraProvider) EthGetUncleByBlockNumberAndIndex()       {}
func (p *InfuraProvider) EthGetUncleCountByBlockHash()            {}
func (p *InfuraProvider) EthGetUncleCountByBlockNumber()          {}
func (p *InfuraProvider) EthGetWork()                             {}
func (p *InfuraProvider) EthHashrate()                            {}
func (p *InfuraProvider) EthMining()                              {}
func (p *InfuraProvider) EthProtocolVersion()                     {}
func (p *InfuraProvider) EthSendRawTransaction()                  {}
func (p *InfuraProvider) EthSubmitWork()                          {}
func (p *InfuraProvider) EthSyncing()                             {}
func (p *InfuraProvider) EthListening()                           {}
func (p *InfuraProvider) EthPeerCount()                           {}
func (p *InfuraProvider) EthVersion()                             {}
func (p *InfuraProvider) EthClientVersion()                       {}
