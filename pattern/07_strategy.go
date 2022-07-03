package main

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
	Паттерн Стратегия предлагает определить семейство схожих алгоритмов, которые часто изменяются или расширяются, и вынести их в собственные классы, называемые стратегиями.
	Плюсы:
	Изолирует код и данные алгоритмов от остальных классов.
	Предоставляет возможность замены одного алгоритма другим в процессе выполнения программы.
	Реализует принцип открытости/закрытости.
	Минусы:
	Требует создание доп. классов
*/

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (Addition) Apply(leftValue, rightValue int) int {
	return leftValue + rightValue
}

type Multiplication struct{}

func (Multiplication) Apply(leftValue, rightValue int) int {
	return leftValue * rightValue
}

func main() {
	add := Operation{Addition{}}
	add.Operate(3, 5)
	mult := Operation{Multiplication{}}
	mult.Operate(3, 5)
}
