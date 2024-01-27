package main

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"snippetbox.heheszlo.com/internal/models"
)

// home is the handler function for the root URL ("/").
// It serves the home page of the web application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "home.tmpl", &templateData{
		Snippets: snippets,
	})
}

// snippetView is the handler function for the "/snippet/view" URL.
// It displays a specific snippet based on an ID provided in the query string.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Handling snippetView for %s", r.URL.Path)
	snippetId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(snippetId)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, http.StatusOK, "view.tmpl", &templateData{
		Snippet: snippet,
	})
}

// snippetCreate is the handler function for the "/snippet/create" URL.
// It handles the creation of new snippets.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST, otherwise return a "Method Not Allowed" error.
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	// If the request method is POST, create a new snippet.

	// dummy data will be removed later
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *application) snippetCreateForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display snippet creation form..."))
}
