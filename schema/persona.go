package schema

// Persona qualifies all information about CV owner
type Persona struct {
	FirstName   string             `yaml:"firstname"`
	LastName    string             `yaml:"lastname"`
	Sex         string             `yaml:"sex"`
	DateOfBirth string             `yaml:"dob"`
	Status      string             `yaml:"status"`
	Picture     string             `yaml:"picture"`
	Nationality string             `yaml:"nationality"`
	Location    *Location          `yaml:"location"`
	Contact     *Contact           `yaml:"contact"`
	Profiles    map[string]Profile `yaml:"profiles"`
}

func (p *Persona) HasProfile(key string) bool {
	_, ok := p.Profiles[key]
	return ok
}
