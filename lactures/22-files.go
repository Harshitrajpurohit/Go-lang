package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "this is my file, i am harshit rajpurohit"
	// create file
	file, err := os.Create("./22newFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content) // both methods can be used to add data in files
	// length, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)

	defer file.Close()

	// read data from the file
	readFile("./22newFile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("data inside file is: ", string(databyte))
}
