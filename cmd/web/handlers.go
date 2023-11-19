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

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with an id %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create new snippet..."))
}
