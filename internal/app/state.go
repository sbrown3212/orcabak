package app

import (
	"github.com/sbrown3212/orcabak/internal/git"
	"github.com/sbrown3212/orcabak/internal/printer"
)

type Config struct {
	SlicerCfgLocation string `mapstructure:"orca-cfg-path"`
	RemoteRepoURL     string `mapstructure:"remote-repo-url"`
}

type State struct {
	AppCfgLocation string
	Config         Config
	Printer        *printer.Printer
	Git            *git.GitCLIClient
}
