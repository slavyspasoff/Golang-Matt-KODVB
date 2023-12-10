package slice_in_depth

import (
	"encoding/json"
	"fmt"
)

func printSlices() {
	var s []int
	t := []int{}
	u := make([]int, 5)
	v := make([]int, 0, 5)

	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil)
}

func sliceToJSON() {
	var a []int
	b := []int{}

	j1, _ := json.Marshal(a) // nil
	j2, _ := json.Marshal(b) // []
	fmt.Printf("%+v %+v\n", string(j1), string(j2))

}

func slicingAndCapacity() {
	arr := [...]int{3, 5, 11, 23, 4}
	sl1 := arr[0:]
	sl2 := arr[0:3]
	sl3 := arr[2:5:5] // [from:to:x] len is to minus from and capacity is to minus from
	fmt.Printf("%+v\nlen:%d\ncap:%d\n", sl3, len(sl3), cap(sl3))
	fmt.Printf("sl[%p] = %v\n", &arr, arr)
	fmt.Printf("a[%p] = %[1]v\n", sl1)
	fmt.Printf("a[%p] = %[1]v\n", sl2)
	fmt.Printf("a[%p] = %[1]v\n", sl3)

}

func Slices() {
	slicingAndCapacity()
}
