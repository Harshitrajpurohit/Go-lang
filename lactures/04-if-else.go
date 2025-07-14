package main

import "fmt"

func main(){
	i := -10

	if i>0{
		fmt.Println("Positive")
	}else if i<0 {
		fmt.Println("negative")
	}else {
		fmt.Println("Zero")
	}

	if i%2==0{
		fmt.Println("even")
	}else {
		fmt.Println("odd")
	}

	var a,b = 10,20

	if a%2==0 && b%2==0 {
		fmt.Println("both are even")
	}

}