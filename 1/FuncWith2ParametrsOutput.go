package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Println(half(x))
	fmt.Println(half(21))
	fmt.Println(half(54))
	fmt.Println(half(100))
}

func half(x int) (int, bool) {
	var isEven, isOdd bool = true, false
	if x%2 == 0 {
		return x / 2, isEven
	} else {
		return x / 2, isOdd
	}
}
