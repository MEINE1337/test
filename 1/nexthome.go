package main

import "fmt"

func swap(x *int, y *int) (int, int) {
	return *y, *x
}
func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}
