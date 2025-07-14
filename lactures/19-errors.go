package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("negative not valid")
	}
	return math.Sqrt(num), nil
}

func main() {
	var num float64 = -10

	var ans, err = Sqrt(num)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
