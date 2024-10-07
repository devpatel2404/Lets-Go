package main

import (
	"html/template"
	"io/fs"
	"modules/internal/models"
	"modules/ui"
	"net/http"
	"path/filepath"
	"time"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	//slices all filepaths that matches the pattern
	pages, err := fs.Glob(ui.Files, "./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		//extracts files names and assigns it to a name variable
		name := filepath.Base(page)

		patterns := []string{
			"html/base.html",
			"html/partials/*.html",
			page,
		}

		ts, err := template.New(name).Funcs(function).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func (app *application) newTemplateDat(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var function = template.FuncMap{
	"humanDate": humanDate,
}
