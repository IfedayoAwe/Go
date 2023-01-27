package main

// Jam jam code that talks about customizing errors, creating error methods to be able to pass a value, exporting the error variables and the type of the error so someone using your functiion can search for an error in the error chain withoutusing string comparisons but using customized error.IS and error.As, creating the unwrap method.
// It's okay to append to a nil slice, read from a nil map, take the length of uninitialized string.

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type errKind int

const (
	_ errKind = iota
	noHeader
	cantReadHeader
	invalidHdrType
	invalidChkLength
	invalidChkType
	invalidLength
)

type WaveError struct {
	kind  errKind
	value int
	err   error
}

func (e WaveError) Error() string {
	var foo string
	switch e.kind {
	case noHeader:
		foo = "no header(file too short?)"

	case cantReadHeader:
		foo = fmt.Sprintf("can't read header [%d]: %s", e.value, e.err.Error())

	case invalidHdrType:
		foo = "invalid header type"

	case invalidChkLength:
		foo = fmt.Sprintf("invalid chunk length: %d", e.value)
	}
	return foo
}

func (e *WaveError) with(val int) *WaveError {
	e1 := e
	e1.value = val
	return e1
}

func (e *WaveError) from(pos int, err error) *WaveError {
	e1 := e
	e1.value = pos
	e1.err = err
	return e1
}

func (w *WaveError) Unwrap() error {
	return w.err
}

func (w *WaveError) Is(t error) bool {
	e, ok := t.(*WaveError)

	if !ok {
		return false
	}

	return e.kind == w.kind
}

var (
	HeaderMissing       = WaveError{kind: noHeader}
	HeaderReadFailed    = WaveError{kind: cantReadHeader}
	InvalidHeaderType   = WaveError{kind: invalidHdrType}
	InvalidChunkkLength = WaveError{kind: invalidChkLength}
	InvalidChunkType    = WaveError{kind: invalidChkType}
	InvalidDataLength   = WaveError{kind: invalidLength}
)

type Header struct {
	TotalLength uint32
}

func DecodeHeader(b []byte) (*Header, []byte, error) {
	var err error
	var pos int
	var val int

	header := Header{TotalLength: uint32(len(b))}
	buf := bytes.NewReader(b)

	if len(b) < 10 {
		return &header, nil, HeaderMissing
	}

	if err = binary.Read(buf, binary.BigEndian, &header); err != nil {
		return &header, nil, HeaderReadFailed.from(pos, err).with(val)
	}

	return &header, b, nil
}
