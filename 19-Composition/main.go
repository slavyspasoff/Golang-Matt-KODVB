package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type Filenamer interface {
	Filename() string
}

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

func Filename(p Pair) string {
	return filepath.Base(p.Path)
}

type PairWithLength struct {
	Pair
	Length int
}

func Embeding() {
	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usr/bin/", "0xdead"}, 133}
	var fn Filenamer = PairWithLength{Pair{"/usr/share/", "0xdebd"}, 174}

	fmt.Println(p)
	fmt.Println(pl)

	fmt.Println(Filename(p))
	fmt.Println(Filename(pl.Pair))

	fmt.Println(fn.Filename())
}

type Fizgig struct {
	*PairWithLength
	Broken bool
}

func PointerEmbeding() {
	fg := Fizgig{
		&PairWithLength{Pair{"/etc", "0xfd11"}, 327},
		false,
	}

	fmt.Println(fg)
}

func main() {
	Embeding()
	PointerEmbeding()
}
