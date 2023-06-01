package main

import (
	"fmt"
	funcdefer "main/funcDefer"
	multiplereturns "main/multiple_returns"
	"main/params"
)

func main() {
	a := map[int]int{4: 1, 7: 2, 8: 3}

	v := params.Do(&a)

	j, k := multiplereturns.MultipleReturn(3, 5)

	fmt.Printf("%+v\n%d\n", a, v)
	fmt.Printf("%d\n%t\n", j, k)

	funcdefer.FuncDefer()
}
