package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"hypergrid-ssn/app"
	"hypergrid-ssn/cmd/hypergrid-ssnd/cmd"
	"hypergrid-ssn/tools"
)

func main() {
	//read config file
	tools.ReadVariablesFromYaml("~/.hypergrid-ssn/config/hypergrid.yaml")

	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
