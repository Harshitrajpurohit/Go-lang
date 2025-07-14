package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums)

	total := 0
	for _, num := range nums { // _, this is to include n-1 value
		total += num
	}
	fmt.Println("total: ", total)
}

func print(num ...int) {
	fmt.Print(num, " ")
	fmt.Println("length is: ", len(num))
}

func main() {

	sum(1, 2, 3, 4, 5)
	sum(1, 2)

	// array
	var a = []int{1, 2, 3, 4}
	print(a...)
	a = append(a, 10)
	print(a...)
	var s = []int{11, 22, 33, 44}
	print(s...)

}
