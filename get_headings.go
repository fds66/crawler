package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	if html == "" {
		return ""
	}
	//fmt.Println("Input", html)
	reader := strings.NewReader(html)
	// look for  text content of <h1> tags if present or  <h2> tag as fallback if neither return empty string
	doc, _ := goquery.NewDocumentFromReader(reader)

	heading := doc.Find("h1").Text()
	if heading != "" {
		//fmt.Println("output: ", heading)
		return heading
	}
	heading = doc.Find("h2").Text()
	if heading != "" {
		//fmt.Println("output: ", heading)
		return heading
	}

	return ""
}
