package main

import (
	"html/template"
	"net/http"
	"path"
)

var pathToTemplates = "./templates"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse the template from disk
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return err
	}
	// execute template passing it data if any
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}
