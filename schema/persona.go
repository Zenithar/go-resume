package schema

import (
	"time"

	"github.com/Sirupsen/logrus"
)

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

func (p *Persona) Age() int {
	date, err := time.Parse("2006-01-02", p.DateOfBirth)
	if err != nil {
		logrus.WithError(err).Fatalf("Unable to parse dateOfBirth as '%s'.", p.DateOfBirth)
	}

	return int(time.Now().Year() - date.Year())
}
