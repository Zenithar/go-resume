package schema

type Company struct {
	Title     string   `yaml:"company"`
	Website   string   `yaml:"website"`
	Activity  string   `yaml:"activity"`
	Sector    string   `yaml:"sector"`
	Location  Location `yaml:"location"`
	Positions []Work   `yaml:"positions"`
}
