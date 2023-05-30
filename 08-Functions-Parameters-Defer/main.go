package main

import "fmt"

func main() {
	var f func(int, int) int

	f = func(x, y int) int {
		return x * y
	}

	fmt.Printf("%d\n", f(9, 3))
}
