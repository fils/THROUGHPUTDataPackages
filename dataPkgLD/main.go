package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"oceanleadership.org/frictionless/THROUGHPUTDataPackages/dataPkgLD/landingpage"
)

// TODO:
// How will schemaFromPackage deal with data in an S3 system?
// How do I materialize the datapackage?

// MyServer struct for mux router
type MyServer struct {
	r *mux.Router
}

func main() {
	log.Println("Make a landing page for a data package")

	searchroute := mux.NewRouter()
	searchroute.HandleFunc("/lptest", landingpage.PresentPackage) // the REAL handler for this URL
	http.Handle("/lptest", searchroute)

	ar := mux.NewRouter()
	ar.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.Handle("/assets/", &MyServer{ar})

	log.Printf("About to listen on 9900. Go to http://127.0.0.1:9900/")

	err := http.ListenAndServe(":9900", nil)
	// http 2.0 http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Let the Gorilla work
	s.r.ServeHTTP(rw, req)
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}
