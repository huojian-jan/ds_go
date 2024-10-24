package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, b := 10, 2

	sum := add(a, b)
	fmt.Println("sum=" + strconv.Itoa(sum))

	subtract := subtract(a, b)
	fmt.Println("sub=" + strconv.Itoa(subtract))

	multiply := multiply(a, b)
	fmt.Println("multiply=" + strconv.Itoa(multiply))

}
