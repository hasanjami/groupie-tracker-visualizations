package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Register specific handlers
	http.HandleFunc("/", serveArtistGalleryHTML)
	http.HandleFunc("/artist/", serveArtistHTML)

	// Serve static files (like images and HTML) from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ArtistDetail struct to hold artist details for the HTML template
type ArtistDetail struct {
	Name         string
	Image        string
	FirstAlbum   string
	Members      []string
	CreationDate int
	Concerts     map[string][]string
}

// Function to serve artist gallery page using a template
func serveArtistGalleryHTML(w http.ResponseWriter, r *http.Request) {
	// Check if the path is exactly "/"
	if r.URL.Path != "/" {
		ErrorHandler(w,r,404)
		return
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("static/basic.html")
	if err != nil {
		ErrorHandler(w,r,500)
		return
	}

	artists, _, _, _, err := getData()
	if err != nil {
		ErrorHandler(w,r,500)
		return
	}

	// Prepare artist data for the template
	data := struct {
		Artists []Artist
	}{
		Artists: artists,
	}

	// Render the template with artist data
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w,r,500)
	}
}

// Function to serve artist detail page using a template
func serveArtistHTML(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("static/artist.html")
	if err != nil {
		ErrorHandler(w,r,500)
		return
	}

	// Extract artist ID from URL path
	path := r.URL.Path
	// Remove the leading "/artist/"
	id := path[len("/artist/"):]

	if id == "" {
		ErrorHandler(w,r,400)
		return
	}

	artists, _, _, relations, err := getData()
	if err != nil {
		ErrorHandler(w,r,500)
		return
	}

	var artistDetail ArtistDetail
	found := false

	for _, artist := range artists {
		if fmt.Sprintf("%d", artist.ID) == id {
			artistDetail = ArtistDetail{
				Name:         artist.Name,
				Image:        artist.Image,
				FirstAlbum:   artist.FirstAlbum,
				Members:      artist.Members,
				CreationDate: artist.CreationDate,
				Concerts:     make(map[string][]string),
			}

			for _, relation := range relations.Index {
				if relation.ID == artist.ID {
					artistDetail.Concerts = relation.Relations
					break
				}
			}
			found = true
			break
		}
	}

	if !found {
		ErrorHandler(w,r,404)
		return
	}

	// Render the template with artistDetail data
	err = tmpl.Execute(w, artistDetail)
	if err != nil {
		ErrorHandler(w,r,500)
	}
}


// ErrorHandler handles errors by rendering an appropriate error page
func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	var errorMessage string
	switch statusCode {
	case http.StatusNotFound:
		errorMessage = "Page not found"
	case http.StatusBadRequest:
		errorMessage = "Bad request"
	case http.StatusInternalServerError:
		errorMessage = "Internal server error"
	case http.StatusMethodNotAllowed:
		errorMessage = "Method not allowed"
	default:
		errorMessage = "Unexpected error"
	}
	data := struct {
		ErrorCode    int
		ErrorMessage string
	}{
		ErrorCode:    statusCode,
		ErrorMessage: errorMessage,
	}
	tmpl, err := template.ParseFiles("static/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Template parse error: %v", err)
		return
	}
	w.WriteHeader(statusCode)
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Template execute error: %v", err)
	}
}

