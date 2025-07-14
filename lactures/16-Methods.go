package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func main() {
	r := rect{width: 10, height: 20}

	area_r := r.area()
	fmt.Println(area_r)

}
