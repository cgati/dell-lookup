package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/dell/test/", testHandler)
	http.HandleFunc("/dell/s/", staticHandler)
	http.ListenAndServe(":8089", nil)
}
