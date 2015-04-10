package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

func validateServiceTags(s []byte) ([]string, error) {
	array := strings.Fields(string(s))
	r := regexp.MustCompile("^[0-9a-zA-Z]{7}$")
	for _, e := range array {
		if r.MatchString(e) == false {
			return []string{}, errors.New("at least one service tag was invalid")
		}
	}
	return array, nil
}

func serveFileFromDir(w http.ResponseWriter, r *http.Request, dir, fileName string) {
	filePath := filepath.Join(dir, fileName)
	http.ServeFile(w, r, filePath)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/dell/s/"):]
	serveFileFromDir(w, r, "../static", file)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Dell Lookup", Body: []byte("Test page!")}
	renderTemplate(w, "main", p)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fromClient, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "couldn't read")
	}
	content, err := getServiceTagsAsJson(fromClient)
	if err != nil {
		fmt.Fprintf(w, "There was an error with a provided service tag: "+err.Error())
		return
	}

	fmt.Fprintf(w, string(content))
}
