package main

import (
	"os"
	"text/template"

	"go.zenithar.org/resume/reader"
	"go.zenithar.org/resume/schema"

	"github.com/Sirupsen/logrus"
)

func main() {
	// Extract resume model from YAML
	model := schema.Resume{}
	err := reader.FromFile("./examples/fr.yaml", &model)
	if err != nil {
		logrus.WithError(err).Fatal("Unable to load resume specification.")
	}

	// Generate output
	tmpl := template.New("output")
	if err != nil {
		logrus.WithError(err).Fatal("Unable to load template.")
	}

	tmpl, err = tmpl.Delims("{~", "~}").ParseGlob("./templates/*.tmpl")
	if err != nil {
		logrus.WithError(err).Fatal("Unable to parse template.")
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "latex.tmpl", &model)
	if err != nil {
		logrus.WithError(err).Fatal("Unable to merge template.")
	}
}
