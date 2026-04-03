package domain

type Config struct {
	OrcaCfgPath   string `mapstructure:"orca-cfg-path" json:"orca-cfg-path,omitempty"`
	RemoteRepoURL string `mapstructure:"remote-repo-url" json:"remote-repo-url,omitempty"`
}
