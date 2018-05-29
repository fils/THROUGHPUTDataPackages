package lphandler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/frictionlessdata/datapackage-go/datapackage"
)

// PresentPackage is the handler for the landing page action
func PresentPackage(w http.ResponseWriter, r *http.Request) {
	so := schemaFromPackage()
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

func schemaFromPackage() string {
	pkg, err := datapackage.Load("../datapackages/simple-geojson/datapackage.json")
	if err != nil {
		log.Printf("Error getting package: %v\n", err)
		return "Error getting package"
	}

	resource := pkg.GetResource("example") // could I also read the "schemaorg" JSON file in and use it?
	rc, err := resource.RawRead()
	defer rc.Close()
	contents, _ := ioutil.ReadAll(rc)
	if err != nil {
		log.Printf("Error getting raw read of entry: %v\n", err)
		return "Error getting schema.org entry"
	}

	return string(contents) // just grab the first "thing" and return it..  we are just testing pipelines...
}
