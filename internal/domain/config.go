package domain

var AllConfigKeys = []string{
	"orca-cfg-path",
	"remote-repo-url",
}

type Config struct {
	OrcaCfgPath   string `mapstructure:"orca-cfg-path" json:"orca-cfg-path,omitempty"`
	RemoteRepoURL string `mapstructure:"remote-repo-url" json:"remote-repo-url,omitempty"`
}

type ConfigItem struct {
	Key   string
	Value string
}

func (c Config) ItemList() []ConfigItem {
	var list []ConfigItem

	if c.OrcaCfgPath != "" {
		item := ConfigItem{Key: "orca-cfg-path", Value: c.OrcaCfgPath}
		list = append(list, item)
	}
	if c.RemoteRepoURL != "" {
		item := ConfigItem{Key: "remote-repo-url", Value: c.RemoteRepoURL}
		list = append(list, item)
	}

	return list
}
