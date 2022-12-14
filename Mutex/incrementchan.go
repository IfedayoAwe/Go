package main

// import (
// 	"fmt"
// 	"sync"
// )

// func do() int {
// 	// channel as a counting semaphore with capacity 1 will do same as a mutex by restricting goroutines access to data
// 	m := make(chan bool, 1)

// 	var n int64
// 	var w sync.WaitGroup

// 	for i := 0; i < 1000; i++ {
// 		w.Add(1)

// 		go func() {
// 			m <- true
// 			n++ // DATA RACE
// 			<-m
// 			w.Done()
// 		}()
// 	}

// 	w.Wait()
// 	return int(n)
// }

// func main() {
// 	fmt.Println(do())
// }
