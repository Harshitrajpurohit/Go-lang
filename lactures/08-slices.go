package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruits = []string{}
	fmt.Println(fruits)

	fruits = append(fruits, "apple", "mango", "orange")
	fmt.Println(fruits)

	fmt.Println(fruits[2:])

	fruits = append(fruits[1:])
	fmt.Println(fruits)

	highScore := make([]int, 4)
	highScore[0] = 20
	highScore[1] = 26
	highScore[2] = 10
	highScore[3] = 7
	fmt.Println(highScore)

	// highScore[4] = 999  // this will give error

	highScore = append(highScore, 600, 700)
	fmt.Println(highScore)
	fmt.Println(len(highScore))

	// sorting
	sort.Ints(highScore)
	fmt.Println(highScore)

	fmt.Println(cap(highScore)) // capacity

}
