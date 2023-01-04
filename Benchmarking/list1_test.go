package main

//making or summing a slice is more efficient than making or summing the linked list

import "testing"

type node struct {
	v int
	t *node
}

func insert(i int, n *node) *node {
	t := &node{i, nil}
	if n != nil {
		n.t = t
	}
	return t
}

func mkList(n int) *node {
	var h, t *node
	h = insert(0, h)
	t = insert(1, h)
	for i := 2; i < n; i++ {
		t = insert(i, t)
	}
	return h
}

func sumList(n *node) (i int) {
	for h := n; h != nil; h = h.t {
		i += h.v
	}
	return
}

func mkSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

func sumSlice(s []int) (i int) {
	for _, v := range s {
		i += v
	}
	return i
}

func BenchmarkList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := mkList(1200)
		sumList(l)
	}
}

func BenchmarkSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := mkSlice(1200)
		sumSlice(l)
	}
}
