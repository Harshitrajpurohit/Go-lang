package main

import "fmt"

func add(a int, b int) int { // for reture define type
	return a + b
}

func add2(a, b int) {
	fmt.Println(a + b)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp

}

func values(a int, b int) (int, int) {
	return b, a
}

func array() [5]int {
	return [5]int{}
}

func append10(sl []string) []string {
	sl = append(sl, "10")
	return sl
}

func main() {

	r := add(1, 2)
	fmt.Println(r)

	add2(3, 4)

	a := 5
	var b int = 10
	swap(&a, &b)
	fmt.Println(a, b)

	a, b = values(a, b)
	fmt.Println(a, b)

	// array
	arr := array()
	fmt.Println(arr)

	var sl = []string{"hello", "h1", "h2", "h3"}
	fmt.Println(sl)
	s := append10(sl)
	fmt.Println(s)
}
