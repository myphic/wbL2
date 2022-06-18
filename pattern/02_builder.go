package main

import "fmt"

/*
	Реализовать паттерн «строитель».
	Используется для построения сложных объектов покомпонентно.
	Плюсы:
	Отделить детали реализации от клиентского кода
	Создание объектов пошагово
	Один и тот же код для создания разных объектов
	Минусы:
	Усложнение кода
	Строитель и его зависимый класс жестко связаны
*/

type ComputerB struct {
	CPU       string
	RAM       int
	Harddrive string
}

type ComputerBuilder interface {
	CPU(val string) ComputerBuilder
	Memory(val int) ComputerBuilder
	HardDrive(val string) ComputerBuilder

	Build() ComputerB
}

type computerBuilder struct {
	cpu string
	ram int
	hd  string
}

func NewComputerBuilder() ComputerBuilder {
	return &computerBuilder{}
}

func (c *computerBuilder) CPU(val string) ComputerBuilder {
	c.cpu = val
	return c
}
func (c *computerBuilder) Memory(val int) ComputerBuilder {
	c.ram = val
	return c
}
func (c *computerBuilder) HardDrive(val string) ComputerBuilder {
	c.hd = val
	return c
}
func (c *computerBuilder) Build() ComputerB {
	return ComputerB{
		CPU:       c.cpu,
		RAM:       c.ram,
		Harddrive: c.hd,
	}
}

//Laptop builder
type laptopBuilder struct {
	computerBuilder
}

func NewLaptopBuilder() ComputerBuilder {
	return (&laptopBuilder{}).CPU("i3").Memory(4).HardDrive("HDD")
}

func (c *laptopBuilder) Build() ComputerB {
	return ComputerB{CPU: c.cpu}
}

// /Laptop builder

// Director

type Director struct {
	c ComputerBuilder
}

func NewDirector(c ComputerBuilder) *Director {
	return &Director{
		c: c,
	}
}

func (d *Director) BuildComputer() ComputerB {
	return d.c.Build()
}

func (d *Director) SetBuilder(c ComputerBuilder) {
	d.c = c

}

func main() {
	cBuilder := NewComputerBuilder()
	cBuilder = cBuilder.CPU("i5").Memory(8).HardDrive("SSD")
	director := NewDirector(cBuilder)
	computer := director.BuildComputer()

	laptopBuilder := NewLaptopBuilder()
	director.SetBuilder(laptopBuilder)
	laptop := director.BuildComputer()
	fmt.Println(computer, laptop)
}
