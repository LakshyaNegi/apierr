package apierr

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

type ArgDef struct {
	Name    string `yaml:"name"`
	ArgType string `yaml:"arg_type"`
}

type ErrorDef struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	ErrType     string   `yaml:"err_type"`
	ErrCode     string   `yaml:"err_code"`
	ErrMsg      string   `yaml:"err_msg"`
	DisplayMsg  string   `yaml:"display_msg"`
	StatusCode  int      `yaml:"status_code"`
	Retryable   bool     `yaml:"retryable"`
	Args        []ArgDef `yaml:"args"`
}

type ErrorsFile struct {
	Errors   []ErrorDef `yaml:"errors"`
	ErrTypes []string   `yaml:"err_types"`
	ErrCodes []string   `yaml:"err_codes"`
}

const errorTemplate = `// Code generated by error-gen. DO NOT EDIT.

package generr

import (
	"apierr"
	"fmt"
)

const (
	// Error Names
	{{- range .Errors}}
	Error{{.Name}} = "{{.Name}}" // {{.Description}}
	{{- end}}

	// Error Types
	{{- range .ErrTypes}}
	ErrType{{.}} = "{{.}}"
	{{- end}}

	// Error Codes
	{{- range .ErrCodes}}
	ErrCode{{.}} = "{{.}}"
	{{- end}}
)

{{range .Errors}}
// {{.Name}}Error represents {{.Description}}.
type {{.Name}}Error struct {
	apierr.CustomError
	{{- range .Args}}
	{{.Name}} {{.ArgType}}
	{{- end}}
}

// New{{.Name}}Error creates a new {{.Name}}Error.
func New{{.Name}}Error(
	{{- range .Args}}{{.Name}} {{.ArgType}},{{- end}}
) *apierr.CustomError {
	return apierr.New(
		{{.StatusCode}},
		fmt.Sprintf(
			"{{.ErrMsg}}",
			{{range .Args}}{{.Name}},{{end}}
		),
		"{{.DisplayMsg}}",
		"{{.ErrType}}",
		"{{.ErrCode}}",
		{{if .Retryable}}true{{else}}false{{end}},
	)
}

{{- $errorName := .Name -}}
{{range .Args}}

// Get{{.Name | title}} returns the value of {{.Name}} for {{$errorName}}Error.
func (e *{{$errorName}}Error) Get{{.Name | title}}() {{.ArgType}} {
	return e.{{.Name}}
}
{{end}}{{end}}
`

// TitleCase returns the input string with the first letter capitalized.
func TitleCase(input string) string {
	if len(input) == 0 {
		return ""
	}
	return strings.ToUpper(string(input[0])) + input[1:]
}

func Generate(inputFile, outputFile string) error {
	// Read YAML file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %w", err)
	}

	var errorsFile ErrorsFile
	if err := yaml.Unmarshal(data, &errorsFile); err != nil {
		return fmt.Errorf("error parsing YAML: %w", err)
	}

	// Validate YAML structure
	if len(errorsFile.Errors) == 0 {
		return fmt.Errorf("no errors defined in YAML file")
	}

	errorsFile.ErrTypes, errorsFile.ErrCodes = uniqueErrTypesAndCodes(errorsFile.Errors)

	// Define template functions.
	funcMap := template.FuncMap{
		"title": TitleCase, // Add TitleCase as "title" for the template.
	}

	// Parse the template with the custom function map.
	tmpl, err := template.New("errors").Funcs(funcMap).Parse(errorTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Create output file
	dir := path.Join(filepath.Dir(outputFile), "generr")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	fname := filepath.Base(outputFile)
	out, err := os.Create(filepath.Join(dir, fname))
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer out.Close()

	// Execute template
	if err := tmpl.Execute(out, errorsFile); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	log.Printf("Error types generated successfully to %s!", outputFile)
	return nil
}

func uniqueErrTypesAndCodes(defs []ErrorDef) (et, ec []string) {
	unique := make(map[string]bool, len(defs))
	for _, def := range defs {
		if len(def.ErrType) > 0 && !unique["et"+def.ErrType] {
			et = append(et, def.ErrType)
			unique["et"+def.ErrType] = true
		}
		if len(def.ErrCode) > 0 && !unique["ec"+def.ErrCode] {
			ec = append(ec, def.ErrCode)
			unique["ec"+def.ErrCode] = true
		}
	}
	return et, ec
}
