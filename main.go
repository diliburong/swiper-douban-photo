package main

import (
	"crawler-douban/parse"
	"fmt"
)

var (
	BaseUrl = "https://www.douban.com/photos/album/1660199003/?start=0"
)

func Start() {
	pages := parse.GetPages(BaseUrl)

	for _, page := range pages {
		fmt.Println(page)
	}
}

func main() {
	Start()

}
