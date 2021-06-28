package main

import (
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

}

func QueryBlock(cli *cli.Client) {
	block, err := cli.QueryBlock(19)
	if err != nil {
		log.Printf("err: %+v", err)
	}
	data := block.Data.Data
	info, err := help.GetTransactionInfoFromData(data[0], false)
	if err != nil {
		log.Printf("err: %+v", err)
	}
	log.Printf("%+v", info)
}
