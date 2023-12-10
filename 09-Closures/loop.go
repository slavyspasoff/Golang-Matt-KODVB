package closures

import "fmt"

func Loop() {
	s := make([]func(), 4)

	for i := 0; i < len(s); i++ {
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}

	for i := 0; i < len(s); i++ {
		s[i]()
	}
}
