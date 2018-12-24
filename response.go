package ethrpc

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Response struct {
	JsonRpc string          `json:"jsonrpc"`
	ID      int64           `json:"id"`
	Result  json.RawMessage `json:"result"`
}

type BlockResponse struct {
	Number       int64    `json:"number"`
	Hash         string   `json:"hash"`
	Nonce        string   `json:"nonce"`
	Miner        string   `json:"miner"`
	Difficulty   int64    `json:"difficulty"`
	Size         int64    `json:"size"`
	GasUsed      int64    `json:"gasUsed"`
	Timestamp    int64    `json:"timestamp"`
	Transactions []string `transactions`
	Uncles       []string `uncles`
}

type BlockResponseHex struct {
	Number       string   `json:"number"`
	Hash         string   `json:"hash"`
	Nonce        string   `json:"nonce"`
	Miner        string   `json:"miner"`
	Difficulty   string   `json:"difficulty"`
	Size         string   `json:"size"`
	GasUsed      string   `json:"gasUsed"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `transactions`
	Uncles       []string `uncles`
}

func (b BlockResponseHex) UnHexed() *BlockResponse {
	return &BlockResponse{
		Number:       HexToInt64(b.Number),
		Hash:         b.Hash,
		Nonce:        b.Nonce,
		Miner:        b.Miner,
		Difficulty:   HexToInt64(b.Difficulty),
		Size:         HexToInt64(b.Size),
		GasUsed:      HexToInt64(b.GasUsed),
		Timestamp:    HexToInt64(b.Timestamp),
		Transactions: b.Transactions,
		Uncles:       b.Uncles,
	}
}

func HexToInt64(hex string) int64 {
	hex = strings.TrimPrefix(hex, "0x")
	i, _ := strconv.ParseInt(hex, 16, 64)
	return i
}
