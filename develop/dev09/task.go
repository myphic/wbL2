package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	writeFile("result.html", resp)
}

func writeFile(filename string, m *http.Response) {
	result, err := os.Create(filename)
	defer result.Close()
	if err != nil {
		log.Print("Error with create file")
	}

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Println("Error with reading body")
	}
	result.WriteString(string(body))
}
