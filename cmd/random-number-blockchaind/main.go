package main

import (
	"os"

	"github.com/ahmetson/random-number-blockchain/app"
	"github.com/ahmetson/random-number-blockchain/cmd/random-number-blockchaind/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
