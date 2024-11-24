package main

import (
	"html/template"
	"net/http"
	"log"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}

	// Render the template
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Welcome to My Website",
		Message: "This is a simple website built with Go!",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home route
	http.HandleFunc("/", homeHandler)

	// Start the server
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
