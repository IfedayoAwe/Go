package main

import (
	"fmt"
	"regexp"
	"strings"
)

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
// 	te := "aba abba abbba"
// 	re := regexp.MustCompile(`b+`)
// 	mm := re.FindAllString(te, -1)
// 	id := re.FindAllStringIndex(te, -1)

// 	fmt.Println(mm)
// 	fmt.Println(id)

// 	for _, d := range id {
// 		fmt.Println(te[d[0]:d[1]])
// 	}
// }

func main() {
	te := "aba abba abbbxa"
	re := regexp.MustCompile("ab+a")
	up := re.ReplaceAllStringFunc(te, strings.ToUpper)

	fmt.Println(up)
}
