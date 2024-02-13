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

var results []struct {
	Name, Url, Text string
}

func Scraper() {
	c := colly.NewCollector()

	if os.Args[1] == "ddg" {
		c.OnHTML(".links_main", func(h *colly.HTMLElement) {
			Name := h.ChildText("h2.result__title")
			Url := h.ChildText("a.result__url")
			Text := h.ChildText("a.result__snippet")

			if Name == "" || Url == "" || Text == "" {
				Name, Url, Text = "nil", "nil", "nil"
			}

			results = append(results, struct {
				Name, Url, Text string
			}{Name, strings.ReplaceAll(Url, `"`, "%22"), Text})

		})
	} else if os.Args[1] == "ggl" {
		c.OnHTML(".snippet", func(h *colly.HTMLElement) {
			Name := h.ChildText("div.title.svelte-xz5zli")
			Url := h.ChildAttr("a.h", "href")
			Text := h.ChildText("div.snippet-description")

			if Name == "" || Url == "" || Text == "" {
				Name, Url, Text = "nil", "nil", "nil"
			}

			results = append(results, struct {
				Name, Url, Text string
			}{Name, strings.ReplaceAll(Url, `"`, "%22"), Text})
		})
	} else if os.Args[1] == "yho" {
		c.OnHTML(".algo-sr", func(h *colly.HTMLElement) {
			Name := h.ChildAttr("a.mxw-100p", "aria-label")
			Url := h.ChildAttr("a.mxw-100p", "href")
			Text := h.ChildText("span.fc-falcon")

			if Name == "" || Url == "" || Text == "" {
				Name, Url, Text = "nil", "nil", "nil"
			}

			results = append(results, struct {
				Name, Url, Text string
			}{Name, strings.ReplaceAll(Url, `"`, "%22"), Text})
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
	} else if os.Args[1] == "yho" {
		var Site string
		Site = fmt.Sprintf("https://search.yahoo.com/search?p=%s", strings.ReplaceAll(os.Args[2], "-", "+"))
		c.Visit(Site)
	}

	for i, result := range results {
		fmt.Printf(config.Bold+"Result #%d\n", i+1)
		fmt.Printf(config.Bold+"Name: "+config.Reset+"%s\n"+
			config.Bold+"Url: "+config.Reset+" https://%s\n"+
			config.Bold+"Description: "+config.Reset+
			"%s\n\n-------------------\n\n",
			result.Name, result.Url, result.Text)
	}
}
