package main

// import (
// 	"bytes"
// 	"fmt"
// )

// func main() {
// 	var wc WriterCloser = NewBufferedWriterCloser()
// 	x := NewBufferedWriterCloser()
// 	x.Write([]byte("Hello my people, hope una stand well?"))
// 	x.Close()
// 	wc.Write([]byte("I like learning folang, though challenging but fun"))
// 	wc.Close()
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type Closer interface {
// 	Close() error
// }

// type WriterCloser interface {
// 	Writer
// 	Closer
// }

// type BufferedWriterCloser struct {
// 	buffer *bytes.Buffer
// }

// func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
// 	n, err := bwc.buffer.Write(data)
// 	if err != nil {
// 		return 0, err
// 	}

// 	v := make([]byte, 8)
// 	for bwc.buffer.Len() > 8 {
// 		_, err := bwc.buffer.Read(v)
// 		if err != nil {
// 			return 0, err
// 		}
// 		_, err = fmt.Println(string(v))
// 		if err != nil {
// 			return 0, nil
// 		}
// 	}

// 	return n, nil
// }

// func (bwc *BufferedWriterCloser) Close() error {
// 	for bwc.buffer.Len() > 0 {
// 		data := bwc.buffer.Next(8)
// 		_, err := fmt.Println(string(data))
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func NewBufferedWriterCloser() *BufferedWriterCloser {
// 	return &BufferedWriterCloser{
// 		buffer: bytes.NewBuffer([]byte{}),
// 	}
// }
