package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

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

//Дефолтная сортировка без ключей
func sortStrings(str []string) []string {
	sort.Strings(str)
	return str
}

//Сортировка по столбку (ключ -k)
func sortByColumn(str []string, k int, delimeter string) []string {
	sort.Slice(str, func(i, j int) bool {
		left := strings.Split(str[i], delimeter)
		right := strings.Split(str[j], delimeter)
		if len(left) < k || len(right) < k {
			log.Fatalln("Incorrect flag -k (k>count of column)")
			return false
		}
		return left[k] < right[k]
	})

	return str
}

//Численная сортировка (ключ -n)
func sortNumeric(str []string) []string {
	sort.Slice(str, func(i, j int) bool {
		left, err := strconv.Atoi(str[i])
		if err != nil {
			log.Fatalf("Error with convert to int: from: %s to: %s", left, err)
		}
		right, err := strconv.Atoi(str[j])
		if err != nil {
			log.Fatalf("Error with convert to int: from: %s to: %s", right, err)
		}
		return left < right
	})
	return str
}

//Сортировка в обратном порядке (ключ -r)
func sortReverse(str []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	return str
}

func sortWithoutDuplicates(str []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range str {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	sortStrings(list)
	return list
}

func main() {
	stringsFromFile := readFile("./develop/dev03/in.txt")
	k := flag.Int("k", 0, "sort by column")
	n := flag.Bool("n", false, "numeric sort")
	r := flag.Bool("r", false, "reverse sort")
	u := flag.Bool("u", false, "without duplicates sort")
	flag.Parse()
	result := make([]string, len(stringsFromFile))
	switch {
	case *k != 0:
		result = sortByColumn(stringsFromFile, *k, " ")
	case *n != false:
		result = sortNumeric(stringsFromFile)
	case *r != false:
		result = sortReverse(stringsFromFile)
	case *u != false:
		result = sortWithoutDuplicates(stringsFromFile)
	default:
		result = sortStrings(stringsFromFile)
	}

	fmt.Println(result)
}
