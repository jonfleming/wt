package main

import (
	"fmt"
	"os"

	"github.com/yourusername/html2text"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: html2text <URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	text, err := html2text.ConvertURL(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(text)
}