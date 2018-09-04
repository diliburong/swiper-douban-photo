package parse

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func SaveImage(Url string) {
	res, err := http.Get(Url)
	if err != nil {
		// fmt.Printf("%d HTTP ERROR:%s", paper.Pid, err)
		return
	}

	defer res.Body.Close()
	//按分辨率目录保存图片
	Dirname := "./tmp/"
	if !IsDirExist(Dirname) {
		os.Mkdir(Dirname, 0755)
		fmt.Printf("dir %s created\n", Dirname)
	}

	//根据URL文件名创建文件
	filename := filepath.Base(Url)
	dst, err := os.Create(Dirname + filename)
	if err != nil {
		// fmt.Println("%d HTTP ERROR:%s", paper.Pid, err)
		return
	}
	defer dst.Close()

	// 写入文件
	io.Copy(dst, res.Body)
}

func IsDirExist(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return p.IsDir()
	}
}
