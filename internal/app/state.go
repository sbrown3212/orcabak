package app

import "github.com/sbrown3212/orcabak/internal/git"

type State struct {
	AppCfgLocation    string
	SlicerCfgLocation string
	Git               *git.GitCLIClient
}
