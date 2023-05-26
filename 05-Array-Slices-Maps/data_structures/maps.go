package data_structures

import "fmt"

func Maps() {
	/*
	   Note: A key to a map can be only values that you can use a comparison operator (== / !=) on i.e.: Strings, Int, Arrays, Booleans. Slices and Maps can't be a key to a map.
	*/

	var m map[string]int // declared but not initialized. nil map can be looked at but you can't assign or it will panic
	p := make(map[string]int)

	a, ok := p["the"] // return 0 because if a key is not declared it will return the default nill value
	fmt.Printf("Was is in the map: %v\n", ok)
	b, ok := m["the"] // same as the above example. Even tho m is not initialized it can be looked up
	fmt.Printf("Was is in the map: %v\n", ok)

	fmt.Printf("a:%d\nb:%d\n", a, b)
	//m["the"] = 1 // This will result in panic because m in not initialized
	m = p      // Will assign p to m as a reference
	m["the"]++ // this is the same as m["the"] = 0 + 1 since not initialized keys are returned as the nil value of the type

	fmt.Printf("key in p:%d\nkey in m:%d\n", p["the"], m["the"])

	l := map[string]int{
		"and": 1,
		"the": 2,
		"or":  3,
	}

	fmt.Printf("%#v\n", l)
}
