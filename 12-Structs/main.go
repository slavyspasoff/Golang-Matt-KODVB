package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

type Album struct {
	title  string
	artist string
	year   int
	copies int
}

func mapOfStructs() {
	c := make(map[string]*Employee)
	c["Lamine"] = &Employee{
		Name:   "Lamine",
		Number: 2,
		Boss:   nil,
		Hired:  time.Now(),
	}

	c["Matt"] = &Employee{
		Name:   "Matt",
		Number: 1,
		Boss:   c["Lamine"],
		Hired:  time.Now(),
	}

	fmt.Printf("%T %+[1]v\n", c["Matt"])
	fmt.Printf("%T %+[1]v\n", c["Lamine"])
	fmt.Printf("%T %[1]v\n", c["Lamine"].Number)
}

func anonymousStruct() {
	var album = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"The White Album",
		"The Beatles",
		1968,
		1000000,
	}

	fmt.Printf("%+v\n", album)
}

func structConversion() {
	type T1 struct {
		X int `json:foo`
	}
	type T2 struct {
		X int `json:bar`
	}
	v1 := T1{1}
	v2 := T2{2}

	v1 = T1(v2)
	fmt.Printf("%T %+[1]v\n", v1)
}

func main() {
	var white Album = Album{
		"The White Album",
		"The Beatles",
		1968,
		1000000,
	}
	var soldAnother1 = func(a Album) {
		a.copies++
		fmt.Printf("%d\n", a.copies)
		fmt.Printf("%d\n", white.copies)
	}
	var soldAnother2 = func(a *Album) {
		a.copies++
		fmt.Printf("%d\n", a.copies)
		fmt.Printf("%d\n", white.copies)
	}

	soldAnother1(white)
	soldAnother2(&white)
}
