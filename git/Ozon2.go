package main

import (
	"fmt"
)

var n int
var a, anew [3]int

func Vvod() int {
	fmt.Println("Vvedite chisla:")
	fmt.Scan(&n)
	if n < 123 || n > 999 {
		fmt.Println("Nepravilniy pin")
		fmt.Scan(&n)
	}
	fmt.Printf("Chislo n: %d", n)
	return n
}

func CreateMassiv(n int) [3]int {
	for i := 0; i < len(a); i++ {
		a[i] = n % 10
		n /= 10
		fmt.Println(a[i])
	}
	fmt.Println("Znachenie do sort", a)
	return a
}

func SortMassiv(a [3]int) [3]int {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
				tmp := a[j]
				a[j] = a[i]
				a[i] = tmp
			}
		}
		anew[i] = a[i]
	}
	fmt.Println("Znachenie posle sortirovki", anew)
	return anew
}

func VarOfNumbers(a [3]int) [3]int {
	SortMassiv(a)
	fmt.Println("vhodnaya A:", anew)
	var b [3]int
	for i := 0; i < len(anew); i++ {
		b[0] = anew[i]
		for j := 0; j < len(anew); j++ {
			b[1] = anew[j]
			for k := 0; k < len(anew); k++ {
				b[2] = anew[k]
				if b[1] == b[0] || b[2] == b[0] || b[1] == b[2] {
					continue
				}
				fmt.Println(b)
			}
		}
	}
	return b
}

func main() {
	Vvod()
	CreateMassiv(n)
	VarOfNumbers(a)
}
