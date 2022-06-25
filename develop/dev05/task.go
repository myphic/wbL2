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

func printLine(lineNum bool, value string, index int) {
	if lineNum {
		fmt.Println("number: ", index+1, "value: ", value)
	} else {
		fmt.Println(value)
	}
}

//Пример запуска программы: go run task.go -A=1 "in.txt" "1"
func main() {
	A := flag.Int("A", 0, "after")
	B := flag.Int("B", 0, "before")
	C := flag.Int("C", 0, "context")
	c := flag.Bool("c", false, "count")
	i := flag.Bool("i", false, "ignore-case")
	v := flag.Bool("v", false, "invert")
	F := flag.Bool("F", false, "fixed")
	n := flag.Bool("n", false, "line num")
	flag.Parse()
	args := flag.Args()
	filename := args[0]
	pattern := args[1]
	stringsFromFile := readFile(filename)

	bottom := 0
	top := 0

	switch {
	case *A != 0:
		bottom = *B
	case *B != 0:
		top = *B
	case *C != 0:
		top, bottom = *C, *C
	}
	counter := 0
	for index, value := range stringsFromFile {
		if *F && value == pattern {
			printLine(*n, value, index)
		}
		if *i {
			value = strings.ToLower(value)
			pattern = strings.ToLower(pattern)
		}
		if *v {
			if !strings.Contains(value, pattern) {
				printLine(*n, value, index)
			}
		}
		if !*v && !*F {
			if strings.Contains(value, pattern) {
				counter++
				if !*c {
					topIndex := index - top
					bottomIndex := index + bottom
					if bottomIndex < len(stringsFromFile) || topIndex < 0 {
						for i := topIndex; i <= bottomIndex; i++ {
							printLine(*n, stringsFromFile[i], index)
						}
					} else {
						printLine(*n, stringsFromFile[index], index)
					}
				}
			}
		}
	}
	if *c {
		fmt.Println(counter)
	}
}
