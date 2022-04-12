package main

import (
	"fmt"
)

func main() {
	month := 3
	fmt.Println(QuarterOf(month))
	fmt.Println(QuarterOf(7))
	fmt.Println(QuarterOf(10))
}

func QuarterOf(month int) int {
	if month > 0 && month <= 3 {
		return 1
	} else if month > 3 && month <= 6 {
		return 2
	} else if month > 6 && month <= 9 {
		return 3
	} else {
		return 4
	}
}
