package schema

import "time"

// Resume is the resume information context holder
type Resume struct {
	Me    Persona `yaml:"info"`
	Quote string  `yaml:"quote"`

	Title        string       `yaml:"title"`
	Summary      string       `yaml:"summary"`
	Expectations Expectations `yaml:"expectations"`

	Companies []Company  `yaml:"companies"`
	Education []School   `yaml:"education"`
	Languages []Language `yaml:"languages"`
	SkillSets []SkillSet `yaml:"skillsets"`

	Extras []string `yaml:"extras"`

	Now      time.Time `yaml:"-"`
	Revision string    `yaml:"-"`
}
