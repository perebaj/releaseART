package interfaces

type Chart struct {
	APIVersion   string `yaml:"apiVersion"`
	Name         string `yaml:"name"`
	Description  string `yaml:"description"`
	Version      string `yaml:"version"`
	Dependencies []struct {
		Name       string `yaml:"name"`
		Version    string `yaml:"version"`
		Repository string `yaml:"repository"`
		Condition  string `yaml:"condition,omitempty"`
	} `yaml:"dependencies"`
}
