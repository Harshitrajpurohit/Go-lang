package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)
	fmt.Println(m)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)

	fmt.Println(m["a"]) // value at "a"

	fmt.Println(len(m)) // length of map

	delete(m, "b") // delete key "b"
	fmt.Println(m)

	_, prs := m["c"]
	fmt.Println(prs)

	// another way

	mp1 := map[string]int{}
	fmt.Println(mp1)

	mp2 := map[string]int{"h": 1, "a": 2, "r": 3}
	fmt.Println(mp2)

	if maps.Equal(mp1, mp2) {
		fmt.Println("mp1 == mp2")
	} else {
		fmt.Println("mp1 != mp2")
	}

}
