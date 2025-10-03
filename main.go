package main

import (
	"github.com/zivivle/go/blockchain"
	"github.com/zivivle/go/cli"
	"github.com/zivivle/go/db"
)

func main() {
	defer db.Close()
	cli.Start()
	blockchain.Blockchain()
}
