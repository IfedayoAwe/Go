package main

// This is a concurrent method using unbuffered channels (pairs)

// import (
// 	"crypto/md5"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"runtime"
// )

// type pair struct {
// 	hash string
// 	path string
// }

// type fileList []string
// type results map[string]fileList

// func hashFile(path string) pair {
// 	file, err := os.Open(path)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer file.Close()

// 	hash := md5.New()

// 	if _, err := io.Copy(hash, file); err != nil {
// 		log.Fatal(err)
// 	}

// 	return pair{fmt.Sprintf("%x", hash.Sum(nil)), path}
// }

// func processFiles(paths <-chan string, pairs chan<- pair, done chan<- bool) {
// 	for path := range paths {
// 		pairs <- hashFile(path)
// 	}

// 	done <- true
// }

// func collectHashes(pairs <-chan pair, result chan<- results) {
// 	hashes := make(results)

// 	for p := range pairs {
// 		hashes[p.hash] = append(hashes[p.hash], p.path)
// 	}

// 	result <- hashes
// }

// func searchTree(dir string, paths chan<- string) error {
// 	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil && err != os.ErrNotExist {
// 			return err
// 		}

// 		if info.Mode().IsRegular() && info.Size() > 0 {
// 			paths <- path
// 		}

// 		return nil
// 	})

// 	return err
// }

// func run(dir string) results {
// 	workers := 2 * runtime.GOMAXPROCS(0)
// 	paths := make(chan string)
// 	// if pairs are buffered it'd keeps the workers working making program faster
// 	pairs := make(chan pair)
// 	done := make(chan bool)
// 	result := make(chan results)

// 	for i := 0; i < workers; i++ {
// 		go processFiles(paths, pairs, done)
// 	}

// 	// we need another goroutine so we don't block here
// 	go collectHashes(pairs, result)

// 	if err := searchTree(dir, paths); err != nil {
// 		return nil
// 	}

// 	// we must close the paths channel so the workers stop
// 	close(paths)

// 	// wait for all the workers to be done
// 	for i := 0; i < workers; i++ {
// 		<-done
// 	}

// 	// by closing the pairs we signal that all hashes have been collected; after all the workers are done
// 	close(pairs)

// 	return <-result
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		log.Fatal("Missing parameter, provide dir name!")
// 	}

// 	if hashes := run(os.Args[1]); hashes != nil {
// 		for hash, files := range hashes {
// 			if len(files) > 1 {
// 				fmt.Println(hash[len(hash)-7:], len(files))

// 				for _, file := range files {
// 					fmt.Println(" ", file)
// 				}
// 			}
// 		}
// 	}
// }
