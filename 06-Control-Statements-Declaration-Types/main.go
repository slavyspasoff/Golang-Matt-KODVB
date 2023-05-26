package main

import "fmt"

func main() {
	a := 4

	switch {
	case a < 4:
		fmt.Println("Lesser than 4")
	case a < 5:
		fmt.Println("Lesser than 5")
		fallthrough
	case a < 6:
		fmt.Println("Lesser than 6")
	default:
		fmt.Println("Nope!")

	}

	switch b := 12; b {
	case 11:
		fmt.Println("Is 11")
	case 12:
		fmt.Println("Is 12")
	default:
		fmt.Println("Again, nope!")
	}

}
