package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	/*	resp, _ := http.Get("https://www.biqugeso.org/biquge_132699/58731581_1.html")
		defer resp.Body.Close()
		n, _ := ioutil.ReadAll(resp.Body)*/

	n, _ := ioutil.ReadFile("index.html")

	/*	fmt.Printf("%s", n)*/
	re := regexp.MustCompile(`<div id="content">\n(.*\n)*.*<div id="content_3">`)
	var r1 []string = []string{""}
	result := re.FindAllStringSubmatch(string(n), -1)
	for i := 0; i < len(result); i++ {
		r1[0] += result[i][0]
	}

	re = regexp.MustCompile(`<p>[\s\S]*</p>`)
	result2 := re.FindAllStringSubmatch(r1[0], -1)

	r1[0] = ""
	for i := 0; i < len(result); i++ {
		r1[0] += result2[i][0]
	}
	r1[0] = strings.Replace(r1[0],`</p><p>`,``,-1)
	r1[0] = strings.Replace(r1[0],`</p>`,``,-1)
	r1[0] = strings.Replace(r1[0],`<p>`,``,-1)
	
	fmt.Println(r1[0])
}
