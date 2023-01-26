package main

import (
	"fmt"
)

func main() {
	fmt.Println(Factorial(0))
}

func Factorial(n int) int {
	//if n == 0 {
	//	return 1
	//}
	counter := 1
	for s := 1; s <= n; s++ {
		counter *= s
	}
	return counter
}
