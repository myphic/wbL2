package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
	Вместо того, чтобы объявлять поведение внутри класса, мы делегируем его внешнему объекту(посетителю).
	В нем должны быть объявлены методы посещения каждого типа.
	Плюсы:
	Упрощается добавление новых операций
	Объединение родственных операции в классе Visitor
	Минусы:
	Нельзя использовать если нельзя расширять существующую иерархию структур
	Затруднено добавление новых классов, поскольку нужно обновлять иерархию посетителя
*/

type Visitor interface {
	VisitRectangle(r *Rectangle)
	VisitCircle(c *Circle)
}

type Shape interface {
	GetType() string
	accept(visitor Visitor)
}

type Rectangle struct {
	a int
	b int
}

func (r *Rectangle) accept(v Visitor) {
	v.VisitRectangle(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.VisitCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Area struct {
	area float64
}

func (a *Area) VisitCircle(c *Circle) {
	fmt.Println("Area for circle")
}

func (a *Area) VisitRectangle(c *Rectangle) {
	fmt.Println("Area for rectangle")
}

func main() {
	circle := &Circle{5}
	rectangle := &Rectangle{2, 3}

	area := &Area{}

	circle.accept(area)
	rectangle.accept(area)
}
