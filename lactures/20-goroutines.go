package main

import (
	"fmt"
	"time"
)

func callMe() {
	fmt.Println("yes?")
}

func main() {
	go callMe()
	time.Sleep(time.Second)
	fmt.Println("hello")
}
