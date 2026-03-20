package git

type Runner interface {
	Run(dir, name string, args ...string) ([]byte, error)
}

type GitCLIClient struct {
	Runner Runner
}

func NewGitCLIclient() *GitCLIClient {
	return &GitCLIClient{
		Runner: &ExecRunner{},
	}
}
