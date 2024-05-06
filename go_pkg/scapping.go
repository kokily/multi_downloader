package scrapping

import (
	"bytes"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func WebScrapping() string {
	doc, err := goquery.NewDocument("http://localhost:3000")

	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer

	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")

		b.WriteString(href + ",")
	})

	return b.String()
}
