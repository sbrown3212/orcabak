package main

import (
	"os"

	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/sbrown3212/orcabak/internal/cli"
	"github.com/sbrown3212/orcabak/internal/git"
	"github.com/sbrown3212/orcabak/internal/printer"
)

func main() {
	printer := printer.NewPrinter(os.Stdout, os.Stderr)
	git := git.NewGitCLIclient()
	state := &app.State{
		Printer: printer,
		Git:     git,
	}

	rootCmd := cli.NewRootCmd(state)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
