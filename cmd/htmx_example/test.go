package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/HarrisonHemstreet/go-ws/internal/model"
	"github.com/HarrisonHemstreet/go-ws/internal/template"
)

// Global template definition
var templates = template.Templates

// handleMain serves a page using a template
func handleMain(w http.ResponseWriter, r *http.Request) {
	tenItems := make([]struct{}, 10)
	partnerVendors := []model.PartnerVendor{
		{
			Created:       time.Now(),
			Edited:        time.Now(),
			VideoLink:     sql.NullString{String: "https://example.com/video1", Valid: true},
			Name:          "Vendor One",
			Description:   "Description for Vendor One",
			ImageLink:     "https://example.com/image1.jpg",
			ThumbnailLink: "https://example.com/thumbnail1.jpg",
			Gallery:       []string{"https://example.com/gallery1/1.jpg", "https://example.com/gallery1/2.jpg"},
			ID:            1,
			Account:       101,
			Featured:      1,
			ContactInfo:   1001,
		},
		{
			Created:       time.Now().Add(-24 * time.Hour), // 1 day ago
			Edited:        time.Now(),
			VideoLink:     sql.NullString{String: "", Valid: false}, // No video link
			Name:          "Vendor Two",
			Description:   "Description for Vendor Two",
			ImageLink:     "https://example.com/image2.jpg",
			ThumbnailLink: "https://example.com/thumbnail2.jpg",
			Gallery:       []string{"https://example.com/gallery2/1.jpg"},
			ID:            2,
			Account:       102,
			Featured:      0,
			ContactInfo:   1002,
		},
	}
	err := templates.ExecuteTemplate(w, "dynamic.html", map[string]interface{}{
		"Items":          tenItems,
		"PartnerVendors": partnerVendors,
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
