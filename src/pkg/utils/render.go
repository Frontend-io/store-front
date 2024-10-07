package utils

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateName string) (err error) {
	path := "./templates/" + templateName
	log.Printf("Rendering template for %s", path)

	parsedTemplate, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	parsedTemplate.Execute(w, nil)
	return nil
}
