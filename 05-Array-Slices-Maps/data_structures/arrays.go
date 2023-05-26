package data_structures

import "fmt"

func Arrays() {
	//Arrays are copied by value
	arr1 := [3]int{3, 5, 7}
	arr2 := [...]int{12, 11, 9, 7} // sized by initializer
	arr3 := arr1                   // Copied by value
	arr3[1] = -5

	fmt.Printf("arr1's len: %d\narr2's len: %d\n", len(arr1), len(arr2))
	fmt.Printf("arr1: %d\narr3: %d\n", arr1[1], arr3[1])
}
