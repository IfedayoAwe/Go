package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func foo() int {

	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		go func() {
			// more efficient than using a mutex for synchronizing increment
			atomic.AddInt64(&n, 1)
			w.Done()
		}()
	}

	w.Wait()
	return int(n)
}

func main() {
	fmt.Println(foo())
}
