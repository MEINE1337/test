package main

import "fmt"

func main() {
	var a [11]int
	for i := 0; i < len(a); i++ {
		fmt.Printf("Vvedite %d chislo\n", i)
		fmt.Scan(&a[i])
		if a[i] > 9 {
			fmt.Println("Vvedi znachenie menshe 10")
			fmt.Scan(&a[i])
		} else {
			continue
		}
	}
	fmt.Println(a)
	code := a[1:4]
	fmt.Println(code)
}
