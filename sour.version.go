package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {
	resp, err := soup.Get("https://www.biqugeso.org/biquge_131248/57142665.html")
	if err != nil {
		os.Exit(1)
	}
	/*n, _ := ioutil.ReadFile("index.html")*/
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "booktxt").FindAll("p")
	for _, link := range links {
		fmt.Println(link.Text())
	}
}
