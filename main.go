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
