package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type response struct {
	Item   string `json:"item"`
	Album  string
	Title  string
	Artist string
}

type respWrapper struct {
	response
}

var j1 = `{
	"item": "album",
	"album": {"title": "A Trip on Trips", "artist": "Lyrx"}
}`

var j2 = `{
	"item": "song",
	"song": {"title": "Isabella", "artist": "Lyrx"}
}`

func (r *respWrapper) UnmarshalJSON(b []byte) (err error) {
	var raw map[string]interface{}

	_ = json.Unmarshal(b, &r.response)
	err = json.Unmarshal(b, &raw)

	switch r.Item {
	case "album":
		album, ok := raw["album"].(map[string]interface{})

		if ok {
			if title, ok := album["title"].(string); ok {
				r.Title, r.Album = title, title
			}

			if artist, ok := album["artist"].(string); ok {
				r.Artist = artist
			}

		}

	case "song":
		song, ok := raw["song"].(map[string]interface{})

		if ok {
			if title, ok := song["title"].(string); ok {
				r.Title = title
			}
			if artist, ok := song["artist"].(string); ok {
				r.Artist = artist
			}
		}
	}

	return err
}

func Decode(st string) error {
	var resp1 respWrapper
	var err error

	if err = json.Unmarshal([]byte(st), &resp1); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp1.response)

	return err
}

func main() {
	Decode(j1)
	Decode(j2)
}
