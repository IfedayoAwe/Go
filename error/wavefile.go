package main

// Jam jam code that talks about customizing errors, creating error methods to be able to pass a value, exporting the error variables and the type of the error so someone using your functiion can search for an error in the error chain withoutusing string comparisons but using customized error.IS and error.As, creating the unwrap method.
// It's okay to append to a nil slice, read from a nil map, take the length of uninitialized string.

// import (
// 	"bytes"
// 	"encoding/binary"
// 	"fmt"
// )

// type errKind int

// const (
// 	_ errKind = iota
// 	noHeader
// 	cantReadHeader
// 	invalidHdrType
// 	invalidChkLength
// 	invalidChkType
// 	invalidLength
// )

// type waveError struct {
// 	kind  errKind
// 	value int
// 	err   error
// }

// func (e waveError) Error() string {
// 	var foo string
// 	switch e.kind {
// 	case noHeader:
// 		foo = "no header(file too short?)"

// 	case cantReadHeader:
// 		foo = fmt.Sprintf("can't read header [%d]: %s", e.value, e.err.Error())

// 	case invalidHdrType:
// 		foo = "invalid header type"

// 	case invalidChkLength:
// 		foo = fmt.Sprintf("invalid chunk length: %d", e.value)
// 	}
// 	return foo
// }

// func (e waveError) with(val int) waveError {
// 	e1 := e
// 	e1.value = val
// 	return e1
// }

// func (e waveError) from(pos int, err error) waveError {
// 	e1 := e
// 	e1.value = pos
// 	e1.err = err
// 	return e1
// }

// func (w *waveError) Unwrap() error {
// 	return w.err
// }

// var (
// 	HeaderMissing       = waveError{kind: noHeader}
// 	HeaderReadFailed    = waveError{kind: cantReadHeader}
// 	InvalidHeaderType   = waveError{kind: invalidHdrType}
// 	InvalidChunkkLength = waveError{kind: invalidChkLength}
// 	InvalidChunkType    = waveError{kind: invalidChkType}
// 	InvalidDataLength   = waveError{kind: invalidLength}
// )

// type Header struct {
// 	TotalLength uint32
// }

// func DecodeHeader(b []byte) (*Header, []byte, error) {
// 	var err error
// 	var pos int

// 	header := Header{TotalLength: uint32(len(b))}
// 	buf := bytes.NewReader(b)

// 	if len(b) < 10 {
// 		return &header, nil, HeaderMissing
// 	}

// 	if err = binary.Read(buf, binary.BigEndian, &header); err != nil {
// 		return &header, nil, HeaderReadFailed.from(pos, err)
// 	}

// 	return &header, b, nil
// }
