package app

import (
	"github.com/sbrown3212/orcabak/internal/domain"
	"github.com/sbrown3212/orcabak/internal/git"
	"github.com/sbrown3212/orcabak/internal/printer"
)

type State struct {
	AppCfgLocation string
	Config         domain.Config
	Printer        *printer.Printer
	Git            *git.GitCLIClient
}
