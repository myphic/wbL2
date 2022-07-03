package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
	Плюсы:
	Уменьшение зависимости между обработчиками и клиентом.
	Принцип единственной ответственности
	Принцип открытости закрытости
	Минусы:
	Запрос может остаться не обработанным
*/

type section interface {
	execute(*task)
	setNext(section)
}
type material struct {
	next section
}

func (m *material) execute(t *task) {
	if t.materialCollected {
		fmt.Println("Material already collected")
		m.next.execute(t)
		return
	}
	fmt.Println("Material section gathering materials")
	t.materialCollected = true
	m.next.execute(t)
}

func (m *material) setNext(next section) {
	m.next = next
}

type assembly struct {
	next section
}

func (a *assembly) execute(t *task) {
	if t.assemblyExecuted {
		fmt.Println("Assembly already done")
		a.next.execute(t)
		return
	}
	fmt.Println("Assembly section assembling")
	t.assemblyExecuted = true
	a.next.execute(t)
}

func (a *assembly) setNext(next section) {
	a.next = next
}

type packaging struct {
	next section
}

func (p *packaging) execute(t *task) {
	if t.packagingExecuted {
		fmt.Println("Packaging already done")
		p.next.execute(t)
		return
	}
	fmt.Println("Packaging section doing packaging")
}

func (p *packaging) setNext(next section) {
	p.next = next
}

type task struct {
	name              string
	materialCollected bool
	assemblyExecuted  bool
	packagingExecuted bool
}

func main() {
	packaging := &packaging{}

	assembly := &assembly{}
	assembly.setNext(packaging)

	material := &material{}
	material.setNext(assembly)

	task := &task{name: "Car"}
	material.execute(task)
}
