package git

func (g *GitCLIClient) Remote(orcaProfileDir string, args ...string) (string, error) {
	args = append([]string{"remote"}, args...)

	output, err := g.Runner.Run(orcaProfileDir, "git", args...)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
