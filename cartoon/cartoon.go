package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// getOne returns the metadata for one comic by number.
func getOne(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't read: %s\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "skipping %d: got %d\n", i, resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid body: %s\n", err)
		os.Exit(-1)
	}

	return body
}

func main() {
	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		// fails  int
		data []byte
	)

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close()
	}

	// the output will be in form of JSON array.
	// so add the brackets before and after

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	// stop if we get 2 404s in a row (get passed #404)
	// fails < 2 as condition

	// get the first 5 comics
	for i := 1; i <= 5; i++ {
		if data = getOne(i); data == nil {
			// fails++
			continue
		}

		if cnt > 0 {
			fmt.Fprint(output, ",") //OB1
		}

		_, err = io.Copy(output, bytes.NewBuffer(data))

		if err != nil {
			fmt.Fprintf(os.Stderr, "stopped: %s\n", err)
			os.Exit(-1)
		}

		// fails = 0
		cnt++
	}

	fmt.Fprintf(os.Stdout, "read %d comics\n", cnt)
}
