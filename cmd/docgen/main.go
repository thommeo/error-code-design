package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/thommeo/error-code-design/pkg/errors"
)

const docTemplate = `# Error Codes Documentation

This document is auto-generated. Do not edit manually.

## Table of Contents

- [Error Code Format](#error-code-format)
{{range .Sections}}- [{{.Title}}](#{{.Title | anchorID}})
{{end}}

## Error Code Format

All error codes follow the format: E<type><data> where:
- E: Fixed prefix identifying this as an error code
- type: Single base-36 character (0-9,A-Z) identifying the error code format
- data: Variable-length base-36 encoded data specific to each format

Base-36 encoding uses digits 0-9 and letters A-Z to pack more information into fewer characters while remaining human-readable.

{{range .Sections}}
## {{.Title}}

{{.Description}}

| {{range .Headers}}{{.}} | {{end}}
|{{range .Headers}}----|{{end}}
{{range .Rows}}| {{range .}}{{.}} | {{end}}
{{end}}

{{end}}`

type DocSection struct {
	Title       string
	Description string
	Headers     []string
	Rows        [][]string
}

type DocData struct {
	Sections []DocSection
}

// getSections returns sections grouped by code type
func getSections() []DocSection {
	var sections []DocSection

	// Get error types
	errorTypes := []errors.ErrorType{
		errors.TinyCode{},
		errors.SimpleCode{},
		errors.Simple511Code{},
		errors.AppComponentErrorCode{},
	}

	// Process each error type
	for _, et := range errorTypes {
		docSection := et.GetDocSection()
		perms := et.GetPermutations()

		// Sort permutations by code
		sort.Slice(perms, func(i, j int) bool {
			return perms[i].Code < perms[j].Code
		})

		// Add rows
		var rows [][]string
		for _, p := range perms {
			rows = append(rows, p.TableFields)
		}

		sections = append(sections, DocSection{
			Title:       docSection.Title,
			Description: docSection.Description,
			Headers:     docSection.Headers,
			Rows:        rows,
		})
	}

	return sections
}

// Custom template function to convert section titles to anchor IDs
func anchorID(title string) string {
	// Simple conversion: lowercase and replace spaces with hyphens
	// In a real implementation, you might want to handle more cases
	// like removing special characters, etc.
	result := ""
	for _, c := range title {
		if c == ' ' {
			result += "-"
		} else {
			result += string(c)
		}
	}
	return strings.ToLower(result)
}

func main() {
	var buf bytes.Buffer

	// Create template with custom function
	tmpl := template.New("doc").Funcs(template.FuncMap{
		"anchorID": anchorID,
	})

	// Parse template
	tmpl = template.Must(tmpl.Parse(docTemplate))

	data := DocData{
		Sections: getSections(),
	}

	if err := tmpl.Execute(&buf, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating documentation: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile("docs/error-codes.md", buf.Bytes(), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing documentation: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Documentation generated successfully")
}
