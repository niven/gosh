package net

import (
	"encoding/json"
	"io"
	"net/http"
)

func BasicAuthGetRequest(url, username, password string) (httpStatusCode int, content []byte, err error) {
	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.SetBasicAuth(username, password)

	response, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer response.Body.Close()

	content, err = io.ReadAll(response.Body)
	if err != nil {
		return 0, nil, err
	}

	return response.StatusCode, content, nil
}

func GetJsonWithHeaders(url string, headers http.Header, result any) (httpStatusCode int, content []byte, err error) {

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	if headers != nil {
		req.Header = headers
	}

	response, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer response.Body.Close()

	content, err = io.ReadAll(response.Body)
	if err != nil {
		return 0, nil, err
	}

	err = json.Unmarshal(content, result)
	if err != nil {
		return 0, content, err
	}

	return response.StatusCode, content, nil
}

func GetJson(url string, result any) (httpStatusCode int, content []byte, err error) {
	return GetJsonWithHeaders(url, nil, result)
}
