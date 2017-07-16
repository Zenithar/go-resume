package schema

type Expectations struct {
	Jobs    []string `yaml:"jobs"`
	Domains []string `yaml:"domains"`
}
