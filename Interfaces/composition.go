package main

import (
	"fmt"
	"path/filepath"
)

type FilenamerIncrementer interface {
	Filename() string
	Increment() int
}

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type PairWithLength struct {
	Pair
	Length int
}

func (p *PairWithLength) Increment() int {
	p.Length++
	return p.Length
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usrs", "0xfdfe"}, 78}
	var fn FilenamerIncrementer = &pl

	fmt.Println(p)
	fmt.Println(pl)
	fmt.Println(fn)
	fmt.Println(p.Filename())
	fmt.Println(pl.Filename())
	fmt.Println(fn.Filename())
	fmt.Println(pl.Increment())
	fmt.Println(pl.Increment())
	fmt.Println(fn.Increment())
	fmt.Println(pl.Length)
}
