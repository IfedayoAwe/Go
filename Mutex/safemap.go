package main

import "sync"

type SafeMap struct {
	sync.Mutex // not safe to copy
	m          map[string]int
}

// so methods take a pointer not a value
func (s *SafeMap) Incr(key string) {
	s.Lock()
	defer s.Unlock()

	// only one goroutine can execute this code at the same time
	s.m[key]++
}
