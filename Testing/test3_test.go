package main

import (
	"go/token"
	"reflect"
	"testing"
)

// Table-driven subtests

type scanTest struct {
	name  string
	input string
	want  []token.Token
}

func (st scanTest) run(t *testing.T) {
	got := []token.Token{}

	if !reflect.DeepEqual(st.want, got) {
		t.Error("E suppose equal na")
	}
}

var scanTests = []scanTest{
	{
		name:  "simple",
		input: "fuck this",
		want:  []token.Token{},
	},
	{
		name:  "hard",
		input: "for real",
		want:  []token.Token{},
	},
	{
		name:  "simple hard",
		input: "amala",
		want:  []token.Token{},
	},
}

func TestScanner(t *testing.T) {
	for _, st := range scanTests {
		t.Run(st.name, st.run)
	}
}
