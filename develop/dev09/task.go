package main

import (
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	resp, err := http.Get(os.Args[1])
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Error with getting site")
	}
	s := writeFile("result.html", resp)
	parseFile(s)
}

func writeFile(filename string, m *http.Response) string {
	result, err := os.Create(filename)
	defer result.Close()
	if err != nil {
		log.Print("Error with create file", err)
	}

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Println("Error with reading body")
	}
	result.WriteString(string(body))
	log.Println("File created:", filename)
	return string(body)
}

func parseFile(s string) {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Println("Error with parse: ", err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		for _, val := range n.Attr {
			if n.Type == html.ElementNode && n.Data == "img" {
				if val.Key == "src" || val.Key == "data-src" {
					nodeHandler(val, "images")
				}
			}
			if n.Type == html.ElementNode && n.Data == "link" {
				if val.Key == "href" {
					nodeHandler(val, "css")
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func nodeHandler(attr html.Attribute, folder string) {
	resp, err := http.Get(os.Args[1] + attr.Val)
	splitUrl := strings.Split(attr.Val, "/")

	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Error with getting site")
	}
	writeFile("assets/"+folder+"/"+splitUrl[len(splitUrl)-1], resp)
}
