package main

import (
	"fmt"
	"os"
)

func main() {
	/* Const in go can only be integer(float, char,etc...), strings and booleans.
	Const in Go is completely immutable because of concurrency.*/
	const (
		i = 1
		f = 3.14
		s = "Hello!"
		c = 'A'
		b = false
	)

	var sum float64
	var n int

	for {
		var val float64
		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}
		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("The average is", sum/(float64(n)))
}
