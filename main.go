package main

// Take's 2 cli arguments old and new and replaces old with new for each line of text
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	if len(os.Args) < 3 {
// 		fmt.Fprint(os.Stderr, "not enough args")
// 		os.Exit(-1)
// 	}

// 	old, new := os.Args[1], os.Args[2]
// 	scan := bufio.NewScanner(os.Stdin)
// 	for scan.Scan() {

// 		s := strings.Split(scan.Text(), old)
// 		t := strings.Join(s, new)

// 		fmt.Println(t)
// 	}
// }

// split text in buffer by words, populate the empty map by the text and values will be the increment, slice of struct to contain the distinct key and values, sort it by bigger value.
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// )

// func main() {
// 	scan := bufio.NewScanner(os.Stdin)
// 	words := make(map[string]int)

// 	scan.Split(bufio.ScanWords)

// 	for scan.Scan() {
// 		words[scan.Text()]++
// 	}

// 	type kv struct {
// 		key string
// 		val int
// 	}

// 	var ss []kv

// 	for k, v := range words {
// 		ss = append(ss, kv{k, v})
// 	}

// 	sort.Slice(ss, func(i, j int) bool {
// 		return ss[i].val > ss[j].val
// 	})

// 	for _, v := range ss[:3] {
// 		fmt.Println(v.key, "appears", v.val, "times")
// 	}
// }

// open, read and close multiple files
// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	for _, fname := range os.Args[1:] {
// 		file, err := os.Open(fname)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			continue
// 		}

// 		if _, err := io.Copy(os.Stdout, file); err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			continue
// 		}

// 		file.Close()
// 	}
// }

// Prints the bytes and length of bytes in a file or files

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	for _, fname := range os.Args[1:] {
// 		file, err := os.Open(fname)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			continue
// 		}

// 		data, err := io.ReadAll(file)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			continue
// 		}
// 		fmt.Println("The file", fname, "has", data)
// 		fmt.Println("The file", fname, "has", len(data), "bytes")
// 	}
// }

// Print the number of lines, words, characters and filename of a file or an array of files and their respective totals
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	var tl, tw, tc int
// 	farray := os.Args[1:]
// 	for _, fname := range farray {
// 		file, err := os.Open(fname)
// 		var lc, wc, cc int

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			continue
// 		}

// 		scan := bufio.NewScanner(file)

// 		for scan.Scan() {
// 			lc++
// 			s := scan.Text()
// 			wc += len(strings.Fields(s))
// 			cc += len(s)

// 		}

// 		tl += lc
// 		tw += wc
// 		tc += cc

// 		fmt.Printf("lines: %d, words: %d, characters: %d, filename: %s\n", lc, wc, cc, fname)
// 		file.Close()

// 	}

// 	if len(farray) > 1 {
// 		fmt.Printf("totallines: %d, totalwords: %d, totalchar: %d", tl, tw, tc)
// 	}

// }

// import (
// 	"fmt"
// )

// func fib() func() int {
// 	a, b := 0, 1

// 	return func() int {
// 		a, b = b, a+b
// 		return b
// 	}
// }

// func main() {
// 	f, g := fib(), fib()
// 	fmt.Println(f(), g())
// }

// Closure
// import (
// 	"fmt"
// )

// func main() {
// 	s := make([]func(), 4)

// 	for i := 0; i < 4; i++ {
// 		i2 := i // closure capture
// 		s[i] = func() {
// 			fmt.Printf("%d @ %p\n", i2, &i2)
// 		}
// 	}

// 	for i := 0; i < 4; i++ {
// 		s[i]()
// 	}
// }
