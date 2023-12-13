package main

import "fmt"

type M interface {
	foo() string
	bar()
}

type P struct {
	value  string
	number int
}

func (p P) foo() string {
	return fmt.Sprintf("This is my value %s", p.value)
}
func (p *P) bar() {
	p.number += 2
}

func main() {
	p := P{
		value:  "100$",
		number: 4,
	}
	var m M = &p
	fmt.Println(m.foo())
}
