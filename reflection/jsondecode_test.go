package main

import "testing"

var p1 = `{
	"item": "album",
	"album": {"title": "A Trip on Trips", "artist": "Lyrx"}
}`

var p2 = `{
	"item": "song",
	"song": {"title": "Isabella", "artist": "Lyrx"}
}`

func TestJson(t *testing.T) {
	var err error
	if err = Decode(p1); err != nil {
		t.Error(err)
	}

	if err = Decode(p2); err != nil {
		t.Error(err)
	}

}
