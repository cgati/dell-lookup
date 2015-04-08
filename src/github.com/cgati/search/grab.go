package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func searchServiceTags(serviceTags []string) (string, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", buildURL(serviceTags), nil)
	if err != nil {
		return "", errors.New("error reaching dell")
	}
	request.Header.Set("apikey", getAPIKey())

	response, err := client.Do(request)
	if err != nil {
		return "", errors.New("HTTP protocol error")
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", errors.New("couldn't parse response body")
		}
		return string(contents), nil
	}
}

func buildURL(serviceTags []string) string {
	baseURL := "https://sandbox.api.dell.com/support/v2/assetinfo/warranty/tags.json"
	paramater := "?svctags="
	var buffer bytes.Buffer

	buffer.WriteString(baseURL)
	buffer.WriteString(paramater)

	for index, st := range serviceTags {
		buffer.WriteString(st)
		if index < len(serviceTags)-1 {
			buffer.WriteString("|")
		}
	}

	return buffer.String()
}
