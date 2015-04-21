package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

func sanatizeDellResult(result []byte) []byte {
	input := string(result)

	re := regexp.MustCompile(":([0-9]*)(,|$)")
	sanatized := re.ReplaceAllString(string(input), `:"$1"$2`)
	return []byte(sanatized)
}

func getServiceTagsAsJson(st []byte) ([]byte, error) {
	array, err := validateServiceTags(st)
	if err != nil {
		return []byte{}, errors.New("a provided service tag was invalid")
	}
	content, err := searchServiceTags(array)
	if err != nil {
		return []byte{}, errors.New("there was a problem looking up your service tags. Please try again.")
	}
	dellAssets, err := getWarrantyInformation(content)
	if err != nil {
		return []byte{}, err
	}
	dellAssets = AddWarrantyStatus(dellAssets)
	jsonAssets, _ := json.Marshal(dellAssets)
	return jsonAssets, nil
}

func searchServiceTags(serviceTags []string) ([]byte, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", buildURL(serviceTags), nil)
	if err != nil {
		return []byte{}, errors.New("error reaching dell")
	}
	request.Header.Set("apikey", getAPIKey())

	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return []byte{}, errors.New("couldn't parse response body")
		}
		contents = sanatizeDellResult(contents)
		return contents, nil
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
