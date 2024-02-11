package cmd

import (
	"fmt"
	"os"

	config "github.com/rendick/tsw/settings"
)

func Switch() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . [ddg|ggl] <search_query>")
		os.Exit(0)
	} else if len(os.Args) == 2 {
		switch os.Args[1] {
		case "--help", "-h":
			config.Help()
		case "ddg":
			Scraper()
		case "ggl":
			Scraper()
		default:
			fmt.Println("Try --help for more information.")

		}
	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		case "ddg":
			Scraper()
		case "ggl":
			Scraper()
		}
	}
}
