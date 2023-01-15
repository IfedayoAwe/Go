package main

// import (
// 	"fmt"
// 	"runtime"
// 	"strconv"
// 	"strings"
// )

// func B() string {
// 	_, file, line, _ := runtime.Caller(1)

// 	fmt.Println(file)

// 	fmt.Println(line)

// 	idx := strings.LastIndexByte(file, '/')

// 	fmt.Println(idx)

// 	return "=> " + file[idx+1:] + ":" + strconv.Itoa(line)
// }

// func A() string {
// 	return B()
// }

// func main() {
// 	fmt.Println(A())
// }
