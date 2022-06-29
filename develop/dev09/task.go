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
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		// ...
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, val := range n.Attr {
				if val.Key == "src" || val.Key == "data-src" {
					resp, err := http.Get(os.Args[1] + val.Val)
					splitUrl := strings.Split(val.Val, "/")

					defer resp.Body.Close()
					if err != nil {
						log.Fatalf("Error with getting site")
					}

					writeFile("assets/images/"+splitUrl[len(splitUrl)-1], resp)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
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
