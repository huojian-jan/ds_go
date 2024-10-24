package main

import (
	"errors"
)

func isEvent(num int) (bool, error) {
	if num%2 == 1 {
		return false, errors.New("error number")
	}
	return true, nil
}

func main() {
	// data := []int{1, 2}
	// for _, v := range data {
	// 	res, err := isEvent(v)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	} else {
	// 		fmt.Println(res)
	// 	}
	// }

	convertion_test2()
}
