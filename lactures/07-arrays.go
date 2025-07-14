package main

import "fmt"

func main(){
	var a[5]int
	var b[5]string

	fmt.Println(a) // [0 0 0 0 0]
	fmt.Println(b) // [    ]

	a[2] = 3;
	fmt.Println(a)

	fmt.Println(len(a)) // length of array

	c := [5]int{1,2,3,4,5}
	fmt.Println(c)

	d := [...]int{1,2,3,4,5}
	fmt.Println(d)

	e := [...]int{100,3:400, 500}
	fmt.Println(e)


	// matrix
	mat := [3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	fmt.Println(mat)

	var mat2 [3][2]int

	fmt.Println(mat2)

	mat2 = [3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	fmt.Println(mat2)

	// print all values
	fmt.Println()
	for i:=0;i<len(mat2);i++{
		for j:=0;j<len(mat2[0]);j++{
			fmt.Print(mat2[i][j]," ");
		}
		fmt.Println()
	}
	fmt.Println()

}