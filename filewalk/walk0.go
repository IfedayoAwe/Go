package main

// import (
// 	"crypto/md5"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"path/filepath"
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

// 	return pair{string(hash.Sum(nil)), path}
// }

// func searchTree(dir string) (results, error) {
// 	hashes := make(results)

// 	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil && err != os.ErrNotExist {
// 			return err
// 		}

// 		if info.Mode().IsRegular() && info.Size() > 0 {
// 			h := hashFile(path)
// 			hashes[h.hash] = append(hashes[h.hash], h.path)
// 		}

// 		return nil
// 	})

// 	return hashes, err
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		log.Fatal("Missing parameter, provide dir name!")
// 	}

// 	hashes, err := searchTree(os.Args[1])

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for hash, files := range hashes {
// 		if len(files) > 1 {
// 			fmt.Println(hash[len(hash)-7:], len(files))

// 			for _, file := range files {
// 				fmt.Println(" ", file)
// 			}
// 		}
// 	}
// }
