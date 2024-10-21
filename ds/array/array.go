package main

import "fmt"

func main(){
	array:=[5]int64{}
	fmt.Println(array)

	array[0]=100
	array[1]=100
	array[2]=100
	array[3]=100
	// array[4]=100

	fmt.Println(array)
	fmt.Println(len(array))
	fmt.Println(cap(array))
	
}
