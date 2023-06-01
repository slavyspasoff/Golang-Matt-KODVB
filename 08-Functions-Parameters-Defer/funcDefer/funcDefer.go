package funcdefer

import (
	"fmt"
)

func FuncDefer() {
	// Defer works on a function and not on a block scope
	a := 10
	// defer copies arguments
	defer fmt.Printf("%d\n", a)
	a = 11
	fmt.Printf("%d\n", a)
}
