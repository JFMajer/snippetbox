package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// home is the handler function for the root URL ("/").
// It serves the home page of the web application.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Define a slice containing the file paths for the base layout and home page templates
	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// Use the template.ParseFiles function to read the template files and store the templates in 'ts'
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template, writing the generated HTML to the http.ResponseWriter.
	// The "base" template is used as the 'layout' template.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// snippetView is the handler function for the "/snippet/view" URL.
// It displays a specific snippet based on an ID provided in the query string.
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with an id %d", id)
}

// snippetCreate is the handler function for the "/snippet/create" URL.
// It handles the creation of new snippets.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST, otherwise return a "Method Not Allowed" error.
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// If the request method is POST, create a new snippet.
	w.Write([]byte("Create new snippet..."))
}
