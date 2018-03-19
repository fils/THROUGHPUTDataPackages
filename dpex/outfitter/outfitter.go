package outfitter

import (
	"io/ioutil"
	"log"

	"github.com/frictionlessdata/datapackage-go/datapackage"
)

// SchemaFromPackage Loads the data package descriptor from the specified URL or file path.
// If path has the ".zip" extension, it will be saved in local filesystem and decompressed before loading.
// TODO: Explore load zip via URL call into S3??
func SchemaFromPackage() string {
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
