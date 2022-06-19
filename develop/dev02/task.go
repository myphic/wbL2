package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackString(str string) (string, error) {
	var result strings.Builder
	var stack rune
	var isLetter bool
	for i, r := range str {
		if unicode.IsDigit(r) && i == 0 {
			return "", fmt.Errorf("error in string: %s", str)
		}
		if unicode.IsLetter(r) {
			if isLetter {
				result.WriteRune(stack)
			}
			if i == utf8.RuneCountInString(str)-1 {
				result.WriteRune(r)
			}
			stack = r
			isLetter = true
		}
		if unicode.IsDigit(r) {
			count, _ := strconv.Atoi(string(r))
			for j := 0; j < count; j++ {
				result.WriteRune(stack)
			}
			isLetter = false
		}
	}
	return result.String(), nil
}

func main() {
	unpackStr, err := UnpackString("a4bc2d5e")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(unpackStr)
}
