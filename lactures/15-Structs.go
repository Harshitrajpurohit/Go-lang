package main

import "fmt"

type person struct {
	name string
	age  int
}

func addPerson(name string, age int) *person {
	p := person{name, age}
	return &p
}

func main() {
	fmt.Println(person{"bob", 10})

	fmt.Println(addPerson("harsit", 10))

	student := addPerson("harsit", 10)
	fmt.Println(student)
	fmt.Println(student.name)

}
