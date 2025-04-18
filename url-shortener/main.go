package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	urlMap := map[string]string{
		"gh":          "https://github.com/ddddami",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
		"/test":       "http://localhost:3000",
	}

	mapHandler := MapHandler(urlMap, mux)

	yamlInput := `
- path: /me
  url: https://github.com/ddddami
- path: /urls
  url: https://github.com/ddddami/gophercises/urls
`

	yamlHandler, err := YAMLHandler([]byte(yamlInput), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :4000")
	err = http.ListenAndServe(":4000", yamlHandler)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
