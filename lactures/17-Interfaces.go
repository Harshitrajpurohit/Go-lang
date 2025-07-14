package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())

}

func main() {
	r := rect{width: 3, height: 4}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	// by interface
	measure(r)
}
