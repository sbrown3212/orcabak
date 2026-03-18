/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/sbrown3212/orcabak/cmd"
	"github.com/sbrown3212/orcabak/internal/app"
)

func main() {
	state := &app.State{}

	rootCmd := cmd.NewRootCmd(state)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
