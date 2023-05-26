package format

import "fmt"

func Format() {
	i := 35712
	fmt.Printf("%d %[1]X %#[1]x\n", i)
	f := 3.14
	fmt.Printf("%f %.2[1]f\n", f)
	s := []int{1, 2, 3}
	a := []rune{'a', 'b', 'c'}
	fmt.Printf("%T\n%[1]v\n%#[1]v\n", s)
	fmt.Printf("%T\n%[1]q\n%#[1]v\n", a)
	m := map[string]int{
		"and": 1,
		"or":  2,
	}

	fmt.Printf("%T\n%[1]v\n%#[1]v\n", m)

}
