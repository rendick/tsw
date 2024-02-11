package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	config "github.com/rendick/tsw/settings"
)

var Site string
var Name, Url, Text string
var Log string

func Scraper() {
	c := colly.NewCollector()

	if os.Args[1] == "ddg" {
		c.OnHTML(".links_main", func(h *colly.HTMLElement) {
			Name := h.ChildText("h2.result__title")
			Url := h.ChildText("a.result__url")
			Text := h.ChildText("a.result__snippet")

			fmt.Printf(config.Bold+"Name: "+config.Reset+"%s\n"+
				config.Bold+"Url: "+config.Reset+" https://%s\n"+
				config.Bold+"Description: "+config.Reset+
				"%s\n\n-------------------\n\n",
				Name,
				strings.ReplaceAll(Url, `"`, "%22"),
				Text)
		})
	} else if os.Args[1] == "ggl" {
		c.OnHTML(".snippet", func(h *colly.HTMLElement) {
			Name := h.ChildText("div.title")
			Url := h.ChildAttr("a.h", "href")
			Text := h.ChildText("div.snippet-description")

			fmt.Printf(config.Bold+"Name: "+config.Reset+"%s\n"+
				config.Bold+"Url: "+config.Reset+"%s\n"+
				config.Bold+"Description: "+config.Reset+
				"%s\n\n-------------------\n\n",
				Name,
				strings.ReplaceAll(Url, `"`, "%22"),
				Text)
		})
	}

	c.OnRequest(func(r *colly.Request) {
		log.Printf("%s\n\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		Log = fmt.Sprintf("Request URL: %s failed with response: %d \nError: %s", r.Request.URL, r, err)
		fmt.Println(Log)
		LogsWriter()
	})

	if os.Args[1] == "ddg" {
		Site = fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", strings.ReplaceAll(os.Args[2], "-", "+"))
		c.Visit(Site)
	} else if os.Args[1] == "ggl" {
		var Site string
		Site = fmt.Sprintf("https://search.brave.com/search?q=%s", strings.ReplaceAll(os.Args[2], "-", "+"))
		c.Visit(Site)
	}
}
