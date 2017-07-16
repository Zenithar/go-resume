package schema

// Contact defines persona contact fields
type Contact struct {
	Mobile  string `yaml:"mobile"`
	Email   string `yaml:"email"`
	Website string `yaml:"website"`
}
