package cmd

import (
	"fmt"
	"os"

	config "github.com/rendick/tsw/settings"
)

func Switch() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "--help", "-h":
			config.Help()
		case "ddg":
			Scraper()
		case "ggl":
			fmt.Println("Google.")

		default:
			fmt.Println("Try --help for more information.")

		}
	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		case "ddg":
			Scraper()
		}
	}
}
