package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")

	// hello
	// world

	myDefer()
	// hello
	// 4 3 2 1 0
	// world
}

func myDefer() {
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, " ")
	}

}
