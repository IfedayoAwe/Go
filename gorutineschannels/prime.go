package main

// import "fmt"

// // not an efficient program, a good example for channels

// func generate(limit int, ch chan<- int) {
// 	for i := 2; i <= limit; i++ {
// 		ch <- i
// 	}

// 	close(ch)
// }

// func filter(src <-chan int, dst chan<- int, prime int) {
// 	for i := range src {
// 		if i%prime != 0 {
// 			dst <- i
// 		}
// 	}

// 	close(dst)
// }

// func sieve(limit int) {
// 	ch := make(chan int)

// 	go generate(limit, ch)

// 	for {
// 		prime, ok := <-ch

// 		if !ok {
// 			break
// 		}

// 		ch1 := make(chan int)

// 		go filter(ch, ch1, prime)

// 		ch = ch1

// 		fmt.Print(prime, " ")
// 	}
// }

// func main() {
// 	sieve(100)
// }
