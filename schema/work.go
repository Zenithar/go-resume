package schema

import (
	"fmt"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

type Work struct {
	Title      string   `yaml:"title"`
	StartDate  string   `yaml:"startDate"`
	EndDate    string   `yaml:"endDate"`
	Trainee    bool     `yaml:"trainee"`
	Highlights []string `yaml:"highlights"`
}

// HasEndDate returns true when end date is specified
func (w *Work) HasEndDate() bool {
	return len(strings.TrimSpace(w.EndDate)) > 0
}

func (w *Work) ParsedStartDate() time.Time {
	date, err := time.Parse("2006-01-02", w.StartDate)
	if err != nil {
		logrus.WithError(err).Fatalf("Unable to parse startDate as '%s'.", w.StartDate)
	}
	return date
}

func (w *Work) ParsedEndDate() time.Time {
	date, err := time.Parse("2006-01-02", w.EndDate)
	if err != nil {
		logrus.WithError(err).Fatalf("Unable to parse endDate as '%s'.", w.EndDate)
	}
	return date
}

func (w *Work) DateRange() string {
	startDate := w.ParsedStartDate()
	if w.HasEndDate() {
		endDate := w.ParsedEndDate()
		return fmt.Sprintf("%s %d - %s %d", startDate.Month(), startDate.Year(), endDate.Month(), endDate.Year())
	}

	return fmt.Sprintf("%s %d - ", startDate.Month(), startDate.Year())
}
