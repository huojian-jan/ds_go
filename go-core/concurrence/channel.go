package main

import (
	"fmt"
	"time"
)

func go_run1() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	//time.Sleep(1 * time.Second)
}

func main() {
	go go_run1()
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
