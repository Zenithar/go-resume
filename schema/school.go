package schema

type School struct {
	School   string    `yaml:"school"`
	Location Location  `yaml:"location"`
	Diplomas []Diploma `yaml:"diplomas"`
}
