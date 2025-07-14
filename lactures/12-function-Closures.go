package main

import "fmt"

func mySeq() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := mySeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := mySeq()
	fmt.Println(newInt())
}
