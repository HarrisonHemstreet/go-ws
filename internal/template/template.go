package template

import (
	"html/template"
	"os"
	"path/filepath"
)

// Templates holds all parsed templates.
var Templates *template.Template

func init() {
	// Initialize the template variable
	Templates = template.New("")

	// Walk the template directory and parse all .html files
	err := filepath.Walk("internal/template", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // return any error encountered
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			// Parse the template file and associate it with the Templates object
			_, err := Templates.ParseFiles(path)
			if err != nil {
				return err // return any error encountered during parsing
			}
		}
		return nil
	})
	if err != nil {
		panic(err) // Handle any errors encountered during walking/parsing
	}
}
