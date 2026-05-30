package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(input string) (string, error) {
	if input == "" {
		return "", nil
	}
	fmt.Println("Input: ", input)
	parsed, err := url.Parse(strings.ToLower(input))
	if err != nil {
		return "", err
	}
	trimmedPath := strings.TrimSuffix(parsed.Path, "/")
	output, err := url.JoinPath(parsed.Host, trimmedPath)

	fmt.Println("Output: ", output)
	return output, nil
}
