package main

import ("fmt"
		"time"
)

func main(){
	i:= 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("more")
	}

	tm := time.Now().Weekday()
	fmt.Println(tm)	
	switch tm{
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend")
	default:
		fmt.Println("It is a weekday")
	}
}