package main

import "fmt"

type inter interface {
	getGreet() string
}

type english struct{}
type spanish struct{}

func main() {
	en := english{}
	sp := spanish{}

	printGreet(en)
	printGreet(sp)
}

func (e english) getGreet() string {
	return "Hi there"
}

func (s spanish) getGreet() string {
	return "Hola there"
}

func printGreet(i inter) {
	fmt.Println(i.getGreet())
}
