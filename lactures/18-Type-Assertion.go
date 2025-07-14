package main

import "fmt"

func main() {
	var s interface{} = "hello"

	fmt.Println(s.(string))

	// type assertion with Ok-idiom

	if s, ok := s.(string); ok {
		fmt.Println(s, "is string")
	} else {
		fmt.Println("not a string")
	}

	var v interface{} = 43
	if v, ok := v.(string); ok {
		fmt.Println(v, "is string")
	} else {
		fmt.Println("not a string")
	}

	var w interface{} = 43
	if w, ok := w.(int); ok {
		fmt.Println(w, "is integer")
	} else {
		fmt.Println("not a integer")
	}
}
