package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: html2text <URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	text, err := ConvertURL(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(text)
}

// ConvertURL fetches the content from the given URL and converts it to plain text
func ConvertURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	body := findBody(doc)
	if body == nil {
		return "", fmt.Errorf("no <body> tag found")
	}

	var text strings.Builder
	extractText(body, &text)

	return strings.TrimSpace(text.String()), nil
}

func findBody(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "body" {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if body := findBody(c); body != nil {
			return body
		}
	}
	return nil
}

func extractText(n *html.Node, text *strings.Builder) {
	if n.Type == html.TextNode {
		text.WriteString(strings.TrimSpace(n.Data))
		text.WriteString(" ")
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractText(c, text)
	}

	if n.Type == html.ElementNode && isBlockElement(n.Data) {
		text.WriteString("\n")
	}
}

func isBlockElement(tag string) bool {
	switch tag {
	case "div", "p", "br", "h1", "h2", "h3", "h4", "h5", "h6", "table", "ul", "ol", "li":
		return true
	}
	return false
}
