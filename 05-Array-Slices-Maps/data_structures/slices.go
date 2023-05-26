package data_structures

import "fmt"

func Slice() {
	var a []int
	fmt.Printf("cap a: %d as %+v\n", cap(a), a)
	a = append(a, 2)
	fmt.Printf("cap a: %d as %+v\n", cap(a), a)
	a = append(a, 7)
	fmt.Printf("cap a: %d as %+v\n", cap(a), a)
	a = append(a, 11)
	fmt.Printf("cap a: %d as %+v\n", cap(a), a)

	var b = []int{1, 3}
	fmt.Printf("%+v\n", b)

	d := make([]int, 5)
	fmt.Printf("%+v\n", d)

	for i := 0; i < 7; i++ {
		val := (i + 1) * 2
		if i >= 5 {
			fmt.Printf("cap[%d]: %d\n", i, cap(d))
			d = append(d, val)
		} else {
			d[i] = val
		}
	}
	fmt.Printf("%+v\n", d)
}
