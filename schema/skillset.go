package schema

type SkillSet struct {
	Category string           `yaml:"category"`
	Skills   map[string]int64 `yaml:"skills"`
}
