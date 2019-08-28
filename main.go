package main

import (
	"log"

	"github.com/shitaibin/fabric-sdk-go-sample/cli"
)

const cfgPath = "./config/config.yaml"

func main() {
	c := cli.New(cfgPath)

	// Install
	if err := c.InstallCC(); err != nil {
		log.Panic(err)
	}

	// Instantiate
	if err := c.InstantiateCC(); err != nil {
		log.Panic(err)
	}

	// Query `init`
	if err := c.QueryCC("init"); err != nil {
		log.Panic(err)
	}

	// Invoke
	if err := c.InvokeCC(); err != nil {
		log.Panic(err)
	}

	// 	Query `a`
	if err := c.QueryCC("a"); err != nil {
		log.Panic(err)
	}
}