package reader

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"go.zenithar.org/resume/schema"
)

var (
	includeStringLen = len("!include ")
)

// FromFile imports resume definition from a file
func FromFile(filePath string, resume *schema.Resume) error {
	workDir, fileName := filepath.Split(filePath)

	// Read original file contents into a byte array
	mainFileBytes, err := readFileOrURL(workDir, fileName)
	if err != nil {
		return fmt.Errorf("Unable to retrieve main file source (Error: %s)", err.Error())
	}

	// Get the contents of the main file
	mainFileBuffer := bytes.NewBuffer(mainFileBytes)

	// Pre-process the original file, following !include directive
	preprocessedContentsBytes, err := preProcess(mainFileBuffer, workDir)
	if err != nil {
		return fmt.Errorf("Error preprocessing Resume file (Error: %s)", err.Error())
	}

	// Parse YAML content
	err = yaml.Unmarshal(preprocessedContentsBytes, resume)
	if err != nil {
		return fmt.Errorf("Error unmarshal preprocessed content (Error: %s)", err.Error())
	}

	// Everything is good
	return nil
}

// read raml file/url
func readFileOrURL(workingDir, fileName string) ([]byte, error) {
	// read from URL if it is an URL, otherwise read from local file.
	if isURL(fileName) {
		return readURL(fileName)
	}
	return readFileContents(workingDir, fileName)
}

func readURL(address string) ([]byte, error) {
	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Reads the contents of a file, returns a bytes buffer
func readFileContents(workingDirectory string, fileName string) ([]byte, error) {

	filePath := filepath.Join(workingDirectory, fileName)

	if fileName == "" {
		return nil, fmt.Errorf("File name cannot be nil: %s", filePath)
	}

	// Read the file
	fileContentsArray, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil,
			fmt.Errorf("Could not read file %s (Error: %s)",
				filePath, err.Error())
	}

	return fileContentsArray, nil
}

// returns true if the path is an HTTP URL
func isURL(path string) bool {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		if _, err := url.Parse(path); err == nil {
			return true
		}
	}
	return false
}

// preProcess acts as a preprocessor for a RAML document in YAML format,
// including files referenced via !include. It returns a pre-processed document.
func preProcess(originalContents io.Reader, workingDirectory string) ([]byte, error) {

	// NOTE: Since YAML doesn't support !include directives, and since go-yaml
	// does NOT play nice with !include tags, this has to be done like this.
	// I am considering modifying go-yaml to add custom handlers for specific
	// tags, to add support for !include, but for now - this method is
	// GoodEnough(TM) and since it will only happen once, I am not prematurely
	// optimizing it.

	var preprocessedContents bytes.Buffer

	// Go over each line, looking for !include tags
	scanner := bufio.NewScanner(originalContents)
	var line string

	// Scan the file until we reach EOF or error out
	for scanner.Scan() {
		line = scanner.Text()

		// Did we find an !include directive to handle?
		if idx := strings.Index(line, "!include"); idx != -1 {

			included := line[idx+includeStringLen:]

			preprocessedContents.Write([]byte(line[:idx]))

			// Get the included file contents
			includedContents, err := readFileOrURL(workingDirectory, included)

			if err != nil {
				return nil,
					fmt.Errorf("Error including file %s:\n    %s",
						included, err.Error())
			}

			// add newline to included content
			prepender := []byte("\n")

			// if it is in response body, we prepend "|" to make it as string
			if strings.HasPrefix(strings.TrimSpace(line), "type") { // in body
				prepender = []byte("|\n")
			}
			includedContents = append(prepender, includedContents...)

			// TODO: Check that you only insert .yaml, .raml, .txt and .md files
			// In case of .raml or .yaml, remove the comments
			// In case of other files, Base64 them first.

			// TODO: Better, step by step checks .. though prolly it'll panic
			// Write text files in the same indentation as the first line
			internalScanner :=
				bufio.NewScanner(bytes.NewBuffer(includedContents))

			// Indent by this much
			firstLine := true
			indentationString := ""

			// Go over each line, write it
			for internalScanner.Scan() {
				internalLine := internalScanner.Text()

				preprocessedContents.WriteString(indentationString)
				if firstLine {
					indentationString = strings.Repeat(" ", idx)
					firstLine = false
				}

				preprocessedContents.WriteString(internalLine)
				preprocessedContents.WriteByte('\n')
			}

		} else {

			// No, just a simple line.. write it
			preprocessedContents.WriteString(line)
			preprocessedContents.WriteByte('\n')
		}
	}

	// Any errors encountered?
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading YAML file: %s", err.Error())
	}
	// Return the preprocessed contents
	return preprocessedContents.Bytes(), nil
}
