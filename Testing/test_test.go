package main

import (
	"reflect"
	"testing"
)

type food struct {
	carbonhydrate string
	protein       string
	fat           string
	fluid         string
	vegetables    string
	vitamin       string
}

type foods []food

func TestEqual(t *testing.T) {
	rich := foods{{carbonhydrate: "rice", protein: "chicken", fat: "butter", fluid: "table water", vegetables: "spinash", vitamin: "apple"}, {carbonhydrate: "spaghetti", protein: "pork", fat: "doughnut", fluid: "condensed milk", vegetables: "eforiro", vitamin: "pine apple"}}

	poor := foods{{carbonhydrate: "tuwo", protein: "beans", fat: "bla", fluid: "satchet water", vegetables: "at all", vitamin: "vitamin c"}}

	if reflect.DeepEqual(rich, poor) {
		t.Fatal("That cant be right")
	}
}
