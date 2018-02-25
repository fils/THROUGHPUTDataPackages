package landingpage

import (
	"log"
	"net/http"

	"github.com/alecthomas/template"
	"oceanleadership.org/frictionless/THROUGHPUTDataPackages/dataPkgLD/osproxy"
)

// PresentPackage is the handler for the landing page action
func PresentPackage(w http.ResponseWriter, r *http.Request) {

	log.Println("Start package to LP process")

	//  Input is some UUID   sha1 likely for me....
	// need; schema JSON-LD, is that all I need to make the LP?

	// get the schema.org
	// load the template...
	// build the page  (web components parse the human display)

	so := osproxy.SchemaFromPackage()

	templateFile := "./templates/landingPage.html"

	ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	err = ht.ExecuteTemplate(w, "Q", so) //substitute fields in the template 't', with values from 'user' and write it out to 'w' which implements io.Writer
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}

	// w.Write([]byte("process lp"))

}
