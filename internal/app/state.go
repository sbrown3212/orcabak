package app

import (
	"github.com/sbrown3212/orcabak/internal/git"
	"github.com/sbrown3212/orcabak/internal/printer"
)

type State struct {
	AppCfgLocation    string
	SlicerCfgLocation string
	Printer           *printer.Printer
	Git               *git.GitCLIClient
}
