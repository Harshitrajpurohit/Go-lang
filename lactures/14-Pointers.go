package main

import "fmt"

func toZero(num int) {
	num = 0
}

func toZeroPointer(num *int) {
	*num = 0
}

func main() {

	num := 10
	toZero(num)
	fmt.Println(num) // 10

	toZeroPointer(&num)
	fmt.Println(num) // 0

}
