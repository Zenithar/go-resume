package schema

// Profile is a "other me" link holder
type Profile struct {
	Username string `yaml:"username"`
	URL      string `yaml:"url"`
}
