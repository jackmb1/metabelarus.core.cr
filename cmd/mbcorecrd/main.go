package main

import (
	"os"

	"github.com/metabelarus/mbcorecr/cmd/mbcorecrd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
