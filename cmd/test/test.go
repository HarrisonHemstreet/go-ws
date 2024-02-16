package main

import (
	"net/http"
	"time"

	"github.com/HarrisonHemstreet/go-ws/internal/template"
)

// Global template definition
var templates = template.Templates

// handleMain serves a page using a template
func handleMain(w http.ResponseWriter, r *http.Request) {
	tenItems := make([]struct{}, 10)
	err := templates.ExecuteTemplate(w, "dynamic.html", map[string]interface{}{
		"Items": tenItems,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// handleUpdateContent dynamically updates content using a template
func handleUpdateContent(w http.ResponseWriter, r *http.Request) {
	// Generate dynamic content, for example, the current server time
	data := struct {
		Time string
	}{
		Time: time.Now().Format(time.RFC1123),
	}

	err := templates.ExecuteTemplate(w, "content-update.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Main route using template
	http.HandleFunc("/", handleMain)

	// Dynamic content update route
	http.HandleFunc("/update-content", handleUpdateContent)

	// Start server
	http.ListenAndServe(":8080", nil)
}
