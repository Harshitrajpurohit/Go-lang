package main

import "fmt"

func fac(num int) int {
	if num == 0 {
		return 1
	}

	return fac(num-1) * num
}

func main() {
	num := 5
	ans := fac(num)
	fmt.Println(ans)
}
