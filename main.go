package main

import (
	"encoding/json"
	"fmt"
)

type boy struct {
	Name string `json:"boy_name"`
	Age  int    `json:"boy_age"`
}

func json_serialize_test1() {
	b := boy{Name: "huojian", Age: 18}
	data, _ := json.Marshal(b)
	fmt.Println(string(data))
}

func main() {
	// interfaceTest1()
	map_all_test()
}
