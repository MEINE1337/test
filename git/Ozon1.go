package main

import (
	"fmt"
	"math"
)

var n int
var a [11]int

func main() {
	Vvod()
	ConvertNumber(n)
}

func Vvod() int {
	var nlow, nhigh float64
	nlow = math.Pow(10, 10)
	nhigh = math.Pow(10, 11)
	fmt.Println("Vvedite nomer telephona:")
	fmt.Scan(&n)
	if n < int(nlow) || n >= int(nhigh) {
		fmt.Println("Nepravilno vveden nomer telephona")
		fmt.Scan(&n)
		fmt.Printf("Znachenie ravno %d\n", n)
	} else {
		fmt.Printf("Znachenie ravno %d\n", n)
	}
	return n
}

func ConvertNumber(n int) []int {
	for i := 0; i < len(a); i++ {
		a[i] = n % 10
		n /= 10
		fmt.Println(a[i])
	}
	fmt.Println(a)
	newMatrix := make([]int, 0, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		newMatrix = append(newMatrix, a[i])
	}
	fmt.Println(newMatrix)
	code := newMatrix[1:4]
	fmt.Println(code)
	return code
}
