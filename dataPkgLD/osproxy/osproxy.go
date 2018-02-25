package osproxy

import (
	"log"

	"github.com/frictionlessdata/datapackage-go/datapackage"
)

func GetObjectByURL(urlloc string) {
	log.Printf("Getting file at URL %s", urlloc)

	// steps
	// parse URL to get the ID needed to get the file
	// this will be the sha value of the package zip

	// read from S3
	// pass file object back to

}

// SchemaFromPackage Loads the data package descriptor from the specified URL or file path.
// If path has the ".zip" extension, it will be saved in local filesystem and decompressed before loading.
// Load zip via URL call into S3??
func SchemaFromPackage() string {
	pkg, err := datapackage.Load("../CarpLake/datapackage.json")
	if err != nil {
		log.Println("error in loading data package json")
	}

	// The following reads the tabular data OK
	resource := pkg.GetResource("data") // could I also read the "schemaorg" JSON file in and use it?
	contents, err := resource.ReadAll()
	if err != nil {
		log.Printf("Error getting schema.org entry: %v\n", err)
		// return nil, err
	}

	return contents[0][0] // just grab the first "thing" and return it..  we are just testing pipelines...
}
