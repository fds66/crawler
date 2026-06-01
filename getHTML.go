package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	if len(rawURL) == 0 {
		return "", fmt.Errorf("No url string provided")
	}
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", "BootCrawler/1.0")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error from request", err.Error())
		return "", err
	}
	defer response.Body.Close()
	fmt.Println("Statuscode: ", response.StatusCode)
	if response.StatusCode >= 400 {
		fmt.Println("Error, status oode of response shows an error", response.StatusCode)
		return "", fmt.Errorf("Error, status oode of response shows an error: %v", response.StatusCode)
	}

	var body []byte
	body, err = io.ReadAll(response.Body)

	//fmt.Println(string([]byte(body)))
	/*
		// print out all the headers in the response
		for k, v := range response.Header {
			fmt.Printf("Header field %q, Value %q\n", k, v)
		}
	*/
	contentType := response.Header.Get("content-type")

	if !strings.Contains(contentType, "text/html") {
		fmt.Println("Response is not html")
		return "", fmt.Errorf("Response not html")
	}
	return string(body), nil
}
