package main

import (
	"flag"
	"html"
	"os"
	"strings"
	"text/template"
	"time"

	"go.zenithar.org/resume/reader"
	"go.zenithar.org/resume/schema"

	"github.com/Sirupsen/logrus"
	"github.com/leekchan/gtf"
)

var (
	yamlPath = flag.String("yaml", "", "Defines the YAML filepath.")
	tmplName = flag.String("tmpl", "", "Defines the template to use.")
)

func init() {
	flag.Parse()
}

func initializeFuncMap() map[string]interface{} {
	funcMap := gtf.GtfTextFuncMap

	// Latex escape
	replacer := strings.NewReplacer(
		"%", "\\%",
		"$", "\\$",
		"{", "\\{",
		"_", "\\_",
		"|", "\\textbar",
		">", "\\textgreater",
		"#", "\\#",
		"&", "\\&",
		"}", "\\}",
		"\\", "\\textbackslash",
		"<", "\\textless",
		"^", "\\textasciicircum{}",
		"~", "\\textasciitilde{}",
		"–", "\\--",
		"—", "\\---",
	)
	funcMap["LatexEscape"] = func(text string) string {
		return replacer.Replace(text)
	}

	// HTML encode
	funcMap["HtmlEscape"] = func(text string) string {
		return html.EscapeString(text)
	}

	return funcMap
}

func main() {
	// Extract resume model from YAML
	model := schema.Resume{}
	err := reader.FromFile(*yamlPath, &model)
	if err != nil {
		logrus.WithError(err).WithField("yaml", *yamlPath).Fatal("Unable to load resume specification.")
	}

	// Generate output
	tmpl := template.New("output")
	if err != nil {
		logrus.WithError(err).Fatal("Unable to load template.")
	}

	tmpl, err = tmpl.Funcs(initializeFuncMap()).Delims("{~", "~}").ParseGlob("./templates/*.tmpl")
	if err != nil {
		logrus.WithError(err).Fatal("Unable to parse template.")
	}

	// Assign inherited values
	model.Now = time.Now()

	err = tmpl.ExecuteTemplate(os.Stdout, *tmplName, &model)
	if err != nil {
		logrus.WithError(err).WithField("tmpl", *tmplName).Fatal("Unable to merge template.")
	}
}
