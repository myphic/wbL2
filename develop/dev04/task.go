package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
	Фукнция проверки строк на анаграммы
	Алгоритм:
	1. Пройтись по первой строке, увеличить количество встречаемости символов
	2. Пройтись по второй строке, уменьшить количество встречаемости символов
	3. Пройтись по первой строке, если какой-то символ не равен 0, то вернуть false
	4. Вернуть true
*/
func isAnagram(s1, s2 string) bool {
	lenS := len(s1)
	lenT := len(s2)

	if lenS != lenT {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {
		anagramMap[string(s1[i])]++
	}

	for i := 0; i < lenT; i++ {
		anagramMap[string(s2[i])]--
	}

	for i := 0; i < lenS; i++ {
		if anagramMap[string(s1[i])] != 0 {
			return false
		}
	}

	return true
}

//Приведение слов к нижнему регистру
func sliceToLower(s []string) {
	for i, v := range s {
		s[i] = strings.ToLower(v)
	}
}

//Сортировка анаграмм
func sortAnagrams(m map[string][]string) {
	for _, value := range m {
		sort.Strings(value)
	}
}

//Функция получения анаграмм
func getAnagrams(str []string) map[string][]string {
	sliceToLower(str)
	result := make(map[string][]string, len(str))
	alreadyInResult := make(map[string]bool, len(str))
	for i := 0; i < len(str)-1; i++ {
		for j := 1; j < len(str); j++ {
			if _, ok := alreadyInResult[str[i]]; !ok && i != j {
				if isAnagram(str[i], str[j]) {
					result[str[i]] = append(result[str[i]], str[j])
					alreadyInResult[str[j]] = true
				}
			}
		}
	}
	sortAnagrams(result)
	return result
}

func main() {
	fmt.Println(getAnagrams([]string{"пяТка", "ТЯПКА", "листок", "слиток", "столик", "тест", "пятак"}))
}
