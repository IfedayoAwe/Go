package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

type ByName struct {
	Organs
}

type ByWeight struct {
	Organs
}

type reverse struct {
	ByName
}

func (s Organs) Len() int {
	return len(s)
}

func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func (r reverse) Less(i, j int) bool {
	return r.ByName.Less(j, i)
}

func main() {
	s := Organs{{"brain", 1340}, {"liver", 1494}, {"spleen", 162}, {"pancreas", 131}, {"heart", 290}}
	fmt.Println("len:", len(s))

	fmt.Println("original:", s)
	for i := 1; i < s.Len(); i++ {
		s.Swap(i-1, i)
	}
	fmt.Println("after swap:", s)

	sort.Sort(ByName{s})
	fmt.Println("ByName:", s)
	sort.Sort(ByWeight{s})
	fmt.Println("ByWeight:", s)
	sort.Sort(reverse{ByName{s}})
	fmt.Println("reverseByName:", s)
	sort.Sort(sort.Reverse(ByWeight{s}))
	fmt.Println("sort.Reverse(ByWeight):", s)
}
