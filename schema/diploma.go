package schema

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Diploma struct {
	Title     string   `yaml:"title"`
	StartDate string   `yaml:"startDate"`
	EndDate   string   `yaml:"endDate"`
	Courses   []string `yaml:"courses"`
}

// HasEndDate returns true when end date is specified
func (w *Diploma) HasEndDate() bool {
	return len(strings.TrimSpace(w.EndDate)) > 0
}

func (w *Diploma) HasStartDate() bool {
	return len(strings.TrimSpace(w.StartDate)) > 0
}

func (w *Diploma) ParsedStartDate() time.Time {
	date, err := time.Parse("2006-01-02", w.StartDate)
	if err != nil {
		logrus.WithError(err).Fatalf("Unable to parse startDate as '%s'.", w.StartDate)
	}
	return date
}

func (w *Diploma) ParsedEndDate() time.Time {
	date, err := time.Parse("2006-01-02", w.EndDate)
	if err != nil {
		logrus.WithError(err).Fatalf("Unable to parse endDate as '%s'.", w.EndDate)
	}
	return date
}

func (d *Diploma) DateRange() string {
	startDate := d.ParsedStartDate()
	if d.HasEndDate() {
		endDate := d.ParsedEndDate()
		return fmt.Sprintf("%s %d - %s %d", startDate.Month(), startDate.Year(), endDate.Month(), endDate.Year())
	}

	return fmt.Sprintf("%s %d - ", startDate.Month(), startDate.Year())
}
