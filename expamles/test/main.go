package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/guoxingx/ethrpc"
)

func main() {
	config := map[string]string{"id": "", "secret": ""}
	file, _ := os.Open("config.json")
	defer file.Close()
	json.NewDecoder(file).Decode(&config)
	ip := ethrpc.NewInfuraProvider(config["id"], config["secret"])
	fmt.Println(ip.EthGetBlockByNumber("latest"))
}
