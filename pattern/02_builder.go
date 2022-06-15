package main

import "fmt"

/*
	Реализовать паттерн «строитель».
	Используется для построения сложных объектов покомпонентно.
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

func main() {
	builder := NewComputerBuilder()
	computer := builder.CPU("i5").Memory(8).HardDrive("SSD").Build()
	fmt.Println(computer)
}
