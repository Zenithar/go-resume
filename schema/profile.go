package schema

// Profile is a "other me" link holder
type Profile struct {
	Network  string `yaml:"network"`
	Username string `yaml:"username"`
	URL      string `yaml:"url"`
}
