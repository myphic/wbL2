package main

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
	Плюсы:
	Реализует принцип открытости/закрытости
	Выделяет код производства объектов в одно место, упрощая поддержку кода
	Избавляет главный класс от привязки к конкретным типам объектов
	Позволяет экономить системные ресурсы путем повторного использования уже созданных объектов вместо порождения новых
	Минусы:
	Может привести к созданию больших параллельных иерархий классов, так как для каждого типа объекта надо создать свой подкласс создателя
*/
type filetype string

const (
	CSV filetype = "CSV"
	PNG filetype = "PNG"
	JS  filetype = "JS"
)

type Creator interface {
	CreateFile(file filetype) File
}

type File interface {
	Use() string
}

type ConcreteCreator struct{}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (p *ConcreteCreator) CreateFile(ftype filetype) File {
	var file File

	switch ftype {
	case CSV:
		file = &Csv{string(ftype)}
	case PNG:
		file = &Png{string(ftype)}
	case JS:
		file = &Js{string(ftype)}
	default:
		log.Fatalln("Unknown Action")
	}

	return file
}

type Csv struct {
	filetype string
}

func (p *Csv) Use() string {
	return p.filetype
}

type Png struct {
	filetype string
}

func (p *Png) Use() string {
	return p.filetype
}

type Js struct {
	filetype string
}

func (p *Js) Use() string {
	return p.filetype
}

func main() {
	creator := NewCreator()
	fmt.Println(creator.CreateFile(CSV).Use())
}
