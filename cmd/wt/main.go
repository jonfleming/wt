package main

import (
	"fmt"
	"os"

	"github.com/jonfleming/wt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: wt <URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	text, err := wt.ConvertURL(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(text)
}