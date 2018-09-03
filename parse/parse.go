package parse

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Page struct {
	Page int
	Url  string
}

type Photo struct {
	Src    string
	Width  int
	Height int
}

var (
	pages []Page
	total int
)

func GetPages(startUrl string) []Page {
	doc, err := goquery.NewDocument(startUrl)
	if err != nil {
		log.Fatal(err)
	}

	pages = append(pages, Page{Page: 1, Url: startUrl})

	nextUrl, _ := doc.Find("#content > div > div.article > div.paginator > span.next > a").Eq(0).Attr("href")

	ParseMovies(doc)

	if nextUrl != "" {
		return GetPages(nextUrl)
	} else {

		fmt.Printf("total=%d", total)
		return pages
	}
}

func ParseMovies(doc *goquery.Document) (photoes []Photo) {
	doc.Find("#content > div > div.article > div.photolst  > div.photo_wrap").Each(func(i int, s *goquery.Selection) {

		currPhoto := s.Find("a img").Eq(0)
		src, _ := currPhoto.Attr("src")
		width, _ := currPhoto.Attr("widht")
		height, _ := currPhoto.Attr("height")

		tWidth, _ := strconv.Atoi(width)
		tHeight, _ := strconv.Atoi(height)

		fmt.Println(src)
		total++

		photo := Photo{
			Src:    src,
			Width:  tWidth,
			Height: tHeight,
		}

		photoes = append(photoes, photo)
	})

	return photoes
}
