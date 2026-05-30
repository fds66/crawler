package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getFirstParagraphFromHTML(html string) string {
	if html == "" {
		return ""
	}
	fmt.Println("Input", html)
	reader := strings.NewReader(html)
	// look for  first <p>, if results don't give what we want search for <main> then first <p> within that, if nothing fallback to first <p>, if nothing return ""
	doc, _ := goquery.NewDocumentFromReader(reader)

	main := doc.Find("main")
	if main.Text() != "" {
		paras := main.Find("p")
		para := paras.First().Text()

		fmt.Println("output: ", para)
		return para
	}
	paras := doc.Find("p")
	para := paras.First().Text()
	if para != "" {
		fmt.Println("output: ", para)
		return para
	}

	return ""
}
