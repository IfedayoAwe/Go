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

// fibonanci number to explain closures
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

// Count words and images in a html document #working with trees in go
// import (
// 	"bytes"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"golang.org/x/net/html"
// )

// var raw = `
// <!DOCTYPE html>
// <html>
//   <body>
//     <h1>My First Heading</h1>
//       <p>My first paragraph.</p>
//       <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
//       <img src="xxx.jpg" width="104" height="142">
//   </body>
// </html>`

// func visit(n *html.Node, words, pics *int) {

// 	if n.Type == html.TextNode {
// 		*words += len(strings.Fields(n.Data))
// 	} else if n.Type == html.ElementNode && n.Data == "img" {
// 		*pics++
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		visit(c, words, pics)
// 	}
// }

// func countWordsAndImages(doc *html.Node) (int, int) {
// 	var words, pics int

// 	visit(doc, &words, &pics)

// 	return words, pics
// }

// func main() {
// 	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "parse failed: %\n", err)
// 		os.Exit(-1)
// 	}

// 	words, pics := countWordsAndImages(doc)

// 	fmt.Printf("%d words and %d images\n", words, pics)
// }

// solving the problem of arrays in loops
// import "fmt"

// func main() {
// 	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
// 	a := [][]byte{}

// 	for _, item := range items {
// 		ar := make([]byte, len(item))
// 		copy(ar, item[:])
// 		a = append(a, ar)
// 		fmt.Println(a)
// 	}
// 	fmt.Println(items)
// 	fmt.Println(a)
// }

// run server and return response hello world and any argument after (/) to browser
// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, world! from %s\n", r.URL.Path[1:])
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }

// Get response from the url above
// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	resp, err := http.Get("http://localhost:8080/" + os.Args[1])

// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(-1)
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		body, err := io.ReadAll(resp.Body)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			os.Exit(-1)
// 		}

// 		fmt.Println(string(body))
// 	}
// }

// net/http
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// const url = "https://jsonplaceholder.typicode.com"

// type todo struct {
// 	ID        int    `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }

// func main() {
// 	resp, err := http.Get(url + "/todos/1")

// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(-1)
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		body, err := io.ReadAll(resp.Body)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			os.Exit(-1)
// 		}

// 		var item todo

// 		err = json.Unmarshal(body, &item)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			os.Exit(-1)
// 		}

// 		fmt.Printf("%#v\n", item)
// 	}
// }
