package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/
type command interface {
	execute()
}

//Ключ зажигания
type ignitionKey struct {
	command command
}

func (i *ignitionKey) press() {
	i.command.execute()
}

//Команда запуска
type onCommand struct {
	engine engine
}

func (c *onCommand) execute() {
	c.engine.on()
}

//Команда отключения
type offCommand struct {
	engine engine
}

func (c *offCommand) execute() {
	c.engine.off()
}

type engine interface {
	on()
	off()
}

//Структура машины
type car struct {
	isRunning bool
}

func (c *car) on() {
	c.isRunning = true
	fmt.Println("The car is started")
}

func (c *car) off() {
	c.isRunning = false
	fmt.Println("The car is turned off")
}

func main() {
	car := &car{}
	onCommand := &onCommand{car}
	offCommand := &offCommand{car}
	onKey := &ignitionKey{onCommand}
	onKey.press()
	onKey = &ignitionKey{offCommand}
	onKey.press()
}
