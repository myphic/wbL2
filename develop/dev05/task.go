package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func readFile(filename string) []string {
	var result []string
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

//Пример запуска программы: go run task.go -A=1 "in.txt" "1"
func main() {
	a := flag.Int("A", 0, "after")
	b := flag.Int("B", 0, "before")
	c := flag.Int("C", 0, "context")
	flag.Parse()
	args := flag.Args()
	filename := args[0]
	pattern := args[1]
	stringsFromFile := readFile(filename)
	//result := make([]string, len(stringsFromFile))
	bottom := 0
	top := 0
	switch {
	case *a != 0:
		bottom = *a
	case *b != 0:
		top = *b
	case *c != 0:
		top, bottom = *c, *c
	}
	for index, value := range stringsFromFile {
		if strings.Contains(value, pattern) {
			topIndex := index - top
			bottomIndex := index + bottom
			if bottomIndex < len(stringsFromFile) || topIndex < 0 {
				for i := topIndex; i <= bottomIndex; i++ {
					fmt.Println(stringsFromFile[i])
				}
			} else {
				fmt.Println(stringsFromFile[index])
			}
		}
	}

}
