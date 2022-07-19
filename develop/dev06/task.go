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
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

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

func fieldsKeyHandler(str []string, f int, delimiter string, s bool) {
	if f > len(str) || f < 1 {
		fmt.Println("Error cut the column")
		os.Exit(1)
	}

	for _, v := range str {
		split := strings.Split(v, delimiter)
		if s {
			if len(split) > 1 {
				fmt.Println(split[f-1])
			}
		} else {
			fmt.Println(split[f-1])
		}
	}
}

func main() {
	f := flag.Int("f", 0, "fields")
	d := flag.String("d", "\t", "delimiter")
	s := flag.Bool("s", false, "separated")
	flag.Parse()
	stringsFromFile := readFile("test.txt")
	if *f != 0 || (*d != "\t" && *f != 0) {
		fieldsKeyHandler(stringsFromFile, *f, *d, *s)
	}
}
