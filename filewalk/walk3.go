package main

// Achieves the best result by making a goroutine for every file and directory and limiting how the goroutines access the disk(i/o) by using channels to restrict and allow access and waitgroups to synchronize processes

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type pair struct {
	hash string
	path string
}

type fileList []string
type results map[string]fileList

func hashFile(path string) pair {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return pair{fmt.Sprintf("%x", hash.Sum(nil)), path}
}

func processFiles(path string, pairs chan<- pair, wg *sync.WaitGroup, limits chan bool) {
	defer wg.Done()

	limits <- true

	defer func() {
		<-limits
	}()

	pairs <- hashFile(path)
}

func collectHashes(pairs <-chan pair, result chan<- results) {
	hashes := make(results)

	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}

	result <- hashes
}

func walkDir(dir string, pairs chan<- pair, wg *sync.WaitGroup, limits chan bool) error {
	defer wg.Done()

	visit := func(path string, info os.FileInfo, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err
		}

		// ignore dir itself to avoid an infinite loop!
		if info.Mode().IsDir() && path != dir {
			wg.Add(1)
			go walkDir(path, pairs, wg, limits)
			return filepath.SkipDir
		}

		if info.Mode().IsRegular() && info.Size() > 0 {
			wg.Add(1)
			go processFiles(path, pairs, wg, limits)
		}

		return nil
	}

	limits <- true

	defer func() {
		<-limits
	}()

	return filepath.Walk(dir, visit)
}

func run(dir string) results {
	workers := 2 * runtime.GOMAXPROCS(0)
	limits := make(chan bool, workers)
	pairs := make(chan pair, workers)
	result := make(chan results)
	wg := new(sync.WaitGroup)

	go collectHashes(pairs, result)

	// multi-threaded walk of the directory tree; we need a waitGroup because we don't know how many to wait for
	wg.Add(1)

	err := walkDir(dir, pairs, wg, limits)

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	close(pairs)

	return <-result
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing parameter, provide dir name!")
	}

	if hashes := run(os.Args[1]); hashes != nil {
		for hash, files := range hashes {
			if len(files) > 1 {
				// we will use just 7 chars like git
				fmt.Println(hash[len(hash)-7:], len(files))

				for _, file := range files {
					fmt.Println("  ", file)
				}
			}
		}
	}
}
