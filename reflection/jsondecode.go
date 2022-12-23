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
		inner, ok := raw["song"].(map[string]interface{})

		if ok {
			if title, ok := inner["title"].(string); ok {
				r.Title = title
			}
			if artist, ok := inner["artist"].(string); ok {
				r.Artist = artist
			}
		}
	}

	return err
}

func main() {
	var resp1, resp2 respWrapper
	var err error

	if err = json.Unmarshal([]byte(j1), &resp1); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp1.response)

	if err = json.Unmarshal([]byte(j2), &resp2); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp2.response)
}
