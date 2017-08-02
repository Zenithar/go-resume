package schema

type SkillSet struct {
	Category I18nLabel       `yaml:"category"`
	Levels   []SkillSetLevel `yaml:"levels"`
}

type I18nLabel map[string]string

type SkillSetLevel struct {
	Level        int64       `yaml:"level"`
	Description  I18nLabel   `yaml:"description"`
	Skills       []I18nLabel `yaml:"skills"`
	Technologies []string    `yaml:"technologies"`
}
