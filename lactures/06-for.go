package main

import "fmt"

func main(){

	i := 0
	for i<5{
		fmt.Print(i," ")
		i=i+1
	}
	fmt.Println()

	for j:=1;j<4;j++ {
		fmt.Print(j," ")
	}
	fmt.Println()

	for i := range 3 {
		fmt.Print(i," ")
	}
	fmt.Println()

	for {
		fmt.Println("loop")
		break
	}

	for n = range 100{
		if n%2==0 {
			fmt.Print(n," ");
		}
	}
	fmt.Println()

}