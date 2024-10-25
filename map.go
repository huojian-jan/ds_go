package main

import (
	"fmt"
	"maps"
	"strconv"
)

func map_test1() {
	score := make(map[string]int)
	score["zhangsan"] = 100
	score["xiaohong"] = 89
	score["wangwu"] = 45

	for k, v := range score {
		info := k + "\t" + strconv.Itoa(v)
		fmt.Println(info)

	}

}

func map_test2() {
	m := map[string]int{
		"zhangsan": 100,
	}

	// for k := range m {
	// 	delete(m, k)
	// }
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// l := len(m)

	m1 := map[string]int{}

	maps.Copy(m1, m)
	m_clone := maps.Clone(m)
	m_ptr := &m

	fmt.Printf("m=%v,ptr=%p\n", m, m_ptr)
	fmt.Printf("mclone=%v,ptr=%p", m_clone, &m_clone)

}

func map_all_test() {
	score := make(map[string]int)
	score["zhangsan"] = 100
	score["xiaohong"] = 89
	score["wangwu"] = 45

	maps.DeleteFunc(score, func(k string, v int) bool {
		return v > 60
	})
	fmt.Println(score)
}
