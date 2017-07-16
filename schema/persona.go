package schema

// Persona qualifies all information about CV owner
type Persona struct {
	FirstName   string    `yaml:"firstname"`
	LastName    string    `yaml:"lastname"`
	Sex         string    `yaml:"sex"`
	DateOfBirth string    `yaml:"dob"`
	Status      string    `yaml:"status"`
	Picture     string    `yaml:"picture"`
	Nationality string    `yaml:"nationality"`
	Location    *Location `yaml:"location"`
	Contact     *Contact  `yaml:"contact"`
	Profiles    []Profile `yaml:"profiles"`
}
