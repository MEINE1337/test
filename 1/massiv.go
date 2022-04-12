package main

import "fmt"

func main() {
	var x int = 1248
	fmt.Println(Massiv(x))
	fmt.Println(Massiv(1028))
}

func Massiv(x int) (a [4]int) {
	for i := 0; i < len(a); i++ {
		a[i] = x % 10
		x /= 10
		fmt.Println(a[i])
	}
	return a
}
