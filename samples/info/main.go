package main

import (
	"encoding/hex"
	"github.com/shitaibin/fabric-sdk-go-sample/cli"
	"github.com/shitaibin/fabric-sdk-go-sample/help"
	"log"
)

const (
	org1CfgPath = "./config/org1sdk-config.yaml"
	org2CfgPath = "./config/org2sdk-config.yaml"
)

var (
	peer0Org1 = "peer0.org1.example.com"
	peer0Org2 = "peer0.org2.example.com"
)

func main() {
	org1Client := cli.New(org1CfgPath, "Org1", "Admin", "User1")
	org2Client := cli.New(org2CfgPath, "Org2", "Admin", "User1")

	defer org1Client.Close()
	defer org2Client.Close()

	QueryBlock(org1Client)
	QueryInfo(org1Client)

}

type BlockInfo struct {
	BlockHash   string
	BlockNumber uint64
	Info        help.TransactionDetail
}

func QueryBlock(cli *cli.Client) {
	block, err := cli.QueryBlock(7)
	if err != nil {
		log.Printf("err: %+v", err)
	}
	data := block.Data.Data
	info, err := help.GetTransactionInfoFromData(data[0], true)
	if err != nil {
		log.Printf("err: %+v", err)
	}
	log.Printf("PreviousHash:%+v", hex.EncodeToString(block.Header.PreviousHash))
	log.Printf("%+v", info)
}

func QueryInfo(cli *cli.Client) {
	info, err := cli.QueryInfo()
	if err != nil {
		log.Printf("err: %+v", err)
	}
	log.Printf("%s", hex.EncodeToString(info.BCI.CurrentBlockHash))
}
