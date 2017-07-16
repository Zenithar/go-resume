package schema

// Resume is the resume information context holder
type Resume struct {
	Me Persona `yaml:"info"`

	Title        string       `yaml:"title"`
	Summary      string       `yaml:"summary"`
	Expectations Expectations `yaml:"expectations"`

	Companies []Company  `yaml:"companies"`
	Education []School   `yaml:"education"`
	SkillSets []SkillSet `yaml:"skillsets"`

	Extras []string `yaml:"extras"`
}
