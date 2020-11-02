package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gophercises/urlshort"
)

var fileName string

func init()  {
	flag.StringVar(&fileName, "file", "paths.yaml", "The path to where yaml path config is located")
	flag.Parse()
}

func main() {
	mux := defaultMux()

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := readFile(fileName)

	if err != nil {
		log.Fatalf("Unable to open file %s, #%v", fileName, err)
	}
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mux)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readFile(filePath string) ([]byte, error) {

	return ioutil.ReadFile(filePath)
}
