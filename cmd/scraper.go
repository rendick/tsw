package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func Scraper() {
	if len(os.Args) == 3 {

		c := colly.NewCollector()

		c.OnHTML(".links_main", func(h *colly.HTMLElement) {
			Name := h.ChildText("h2.result__title")
			Url := h.ChildText("a.result__url")
			Text := h.ChildText("a.result__snippet")

			Url = strings.TrimSpace(strings.ReplaceAll(Url, `"`, "%22"))
			Text = strings.TrimSpace(Text)

			fmt.Printf("Name: %s\nUrl: https://%s\nDescription: %s\n\n-------------------\n\n", Name, Url, Text)
		})

		c.OnRequest(func(r *colly.Request) {
			log.Printf("%s\n", r.URL)
		})
		var Site string
		Site = fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", strings.ReplaceAll(os.Args[2], "-", "+"))
		c.Visit(Site)
	} else {
		fmt.Println("Error.")
	}
}
