package closures

import "fmt"

func Fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func PrintFib() {
	f := Fib()
	x := f()
	for ; x < 100; x = f() {
		fmt.Println(x)

	}
}
