package main

import (
	"fmt"
	"main/params"
)

func main() {
	a := []int{1, 2, 3}
	v := params.Do(a)
	fmt.Printf("%+v\n%d\n", a, v)
}
