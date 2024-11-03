package main

import "fmt"

func int_float64(i int) float64 {
	return float64(i)
}

func float64_int(f float64) int {
	return int(f)
}

func convertion_test1() {
	i := 100
	f := int_float64(i)
	fmt.Println(f)
}

func convertion_test2() {
	f := 3.14
	i := float64_int(f)
	fmt.Println(i)
}
