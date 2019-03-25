package main

import "fmt"

type ClassData struct {
	className        string
	fields           map[string]string
	startCalculation bool
}

func (c *ClassData) InitFields() {
	c.fields = make(map[string]string)
}

func (c *ClassData) GetGoType() string {

	str := fmt.Sprintln("package test")

	str += fmt.Sprintf("type %s struct { \n", c.className)

	for a := range c.fields {
		str += fmt.Sprintf("\t %s %s, \n", a, c.fields[a])
	}
	str += "}"

	return str
}
