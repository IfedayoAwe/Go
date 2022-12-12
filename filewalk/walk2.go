package main

// import (
// 	"crypto/md5"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"runtime"
// 	"sync"
// )

// // Faster than walk1, creating a goroutine for every directory and using waitgroups to coordinate their processes
// // Can be made faster by buffering pairs and increasing the number of workers

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

// func walkDir(dir string, paths chan<- string, wg *sync.WaitGroup) error {
// 	defer wg.Done()

// 	visit := func(path string, info os.FileInfo, err error) error {
// 		if err != nil && err != os.ErrNotExist {
// 			return err
// 		}

// 		// ignore dir itself to avoid an infinite loop and skipDir cause a new goroutine now handeles it!
// 		if info.Mode().IsDir() && path != dir {
// 			wg.Add(1)
// 			go walkDir(path, paths, wg)
// 			return filepath.SkipDir
// 		}

// 		if info.Mode().IsRegular() && info.Size() > 0 {
// 			paths <- path
// 		}

// 		return nil
// 	}

// 	return filepath.Walk(dir, visit)
// }

// func run(dir string) results {
// 	workers := 2 * runtime.GOMAXPROCS(0)
// 	paths := make(chan string)
// 	pairs := make(chan pair)
// 	done := make(chan bool)
// 	result := make(chan results)
// 	wg := new(sync.WaitGroup)

// 	for i := 0; i < workers; i++ {
// 		go processFiles(paths, pairs, done)
// 	}

// 	go collectHashes(pairs, result)

// 	wg.Add(1)

// 	err := walkDir(dir, paths, wg)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	wg.Wait()
// 	close(paths)

// 	for i := 0; i < workers; i++ {
// 		<-done
// 	}

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
