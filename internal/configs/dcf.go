package configs

type Dcf struct {
	Metadata     Metadata       `yaml:"metadata"`
	Dependencies []Dependencies `yaml:"dependencies"`
	Terradep     Terradep       `yaml:"terradep"`
}
type Metadata struct {
	Name    string  `yaml:"name"`
	Version float64 `yaml:"version"`
}
type Dependencies struct {
	Name    string  `yaml:"name"`
	Version float64 `yaml:"version"`
	Use     string  `yaml:"use"`
	Branch  string  `yaml:"branch"`
	Repo    string  `yaml:"repo"`
	Subdir  string  `yaml:"subdir"`
	Type    string  `yaml:"type"`
}
