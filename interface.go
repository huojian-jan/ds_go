package main

import "fmt"

type eater interface {
	eat(food string)
}

type man struct {
	name string
}

func (m man) eat(food string) {
	fmt.Println("my name is "+m.name, "i eat "+food)
}

func interfaceTest1() {
	m := man{name: "zhangsan"}
	m.eat("apple")
}

// func main() {
// 	test1()
// }
