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

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type PairWithLength struct {
	Pair
	Length int
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usr", "0xfdfe"}, 78}

	fmt.Println(p)
	fmt.Println(pl)
	fmt.Println(pl.Filename())
}
