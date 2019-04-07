package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gomarkdown/markdown"
	"jaytaylor.com/html2text"
)

var route string

func main() {
	service := os.Args[1]
	repository := os.Args[2]

	if service == "github" {
		route = "https://raw.githubusercontent.com/" + repository + "/master/README.md"
	} else {
		route = "https://gitlab.com/" + repository + "/raw/master/README.md"
	}
	resp, err := http.Get(route)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	md := []byte(body)

	html := string(markdown.ToHTML(md, nil, nil))
	text, err := html2text.FromString(html, html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}
