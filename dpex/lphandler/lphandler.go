package lphandler

import (
	"log"
	"net/http"

	"github.com/alecthomas/template"
	"oceanleadership.org/frictionless/THROUGHPUTDataPackages/dpex/outfitter"
)

// PresentPackage is the handler for the landing page action
func PresentPackage(w http.ResponseWriter, r *http.Request) {
	so := outfitter.SchemaFromPackage()
	templateFile := "./templates/landingPage.html"

	ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	err = ht.ExecuteTemplate(w, "Q", so) //substitute fields in the template 't', with values from 'user' and write it out to 'w' which implements io.Writer
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}
}
