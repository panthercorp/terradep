package configs

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Terradep struct {
	Repositories Repositories `yaml:"repositories"`
}
type Git struct {
	URL                string `yaml:"url"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	Token              string `yaml:"token"`
	SSHPrivateKeysPath string `yaml:"sshPrivateKeysPath"`
	Alias              string `yaml:"alias"`
}
type Headers struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
type HTTP struct {
	URL      string    `yaml:"url"`
	Method   string    `yaml:"method"`
	Username string    `yaml:"username"`
	Password string    `yaml:"password"`
	Headers  []Headers `yaml:"headers"`
	Alias    string    `yaml:"alias"`
}
type Repositories struct {
	Git  []Git  `yaml:"git"`
	HTTP []HTTP `yaml:"http"`
}

func UnmarshallTerradep(file []byte) (Terradep, error) {
	var data Terradep
	if err := yaml.Unmarshal(file, &data); err != nil {
		return Terradep{}, err
	}
	return data, nil
}

func (t *Terradep) GetGitRepository(alias string) (Git, error) {
	for _, r := range t.Repositories.Git {
		if r.Alias == alias {
			return r, nil
		}
	}
	return Git{}, fmt.Errorf("repository not found")
}

func (t *Terradep) GetHTTPRepository(alias string) (HTTP, error) {
	for _, r := range t.Repositories.HTTP {
		if r.Alias == alias {
			return r, nil
		}
	}
	return HTTP{}, fmt.Errorf("repository not found")
}
