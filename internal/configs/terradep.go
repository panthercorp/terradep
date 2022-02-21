package configs

type Terradep struct {
	Repositories Repositories `yaml:"repositories"`
}
type Git struct {
	URL                string `yaml:"url"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
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
