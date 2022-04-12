package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["Max"] = 40
	ages["Sofia"] = 25

	for _, value := range ages {
		fmt.Printf(" %d\n", value)
	}
}
