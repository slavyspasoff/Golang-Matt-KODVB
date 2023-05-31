package declare

import "fmt"

func Declare() {
	var fn func(int, int) bool
	fn = func(a, b int) bool {
		return a < b
	}

	fn2 := func(x, y int) int {
		return x * y
	}

	fmt.Printf("%t\n%d\n", fn(9, 27), fn2(9, 3))
}
