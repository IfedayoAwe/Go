package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32 // Never use float type for money, use a package

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

// Fake database
type database struct {
	mu   sync.Mutex
	data map[string]dollars
}

// // add the handlers
func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for item, price := range db.data {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) fetch(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	item := req.URL.Query().Get("item")

	if _, ok := db.data[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404

		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "item %s has price %s\n", item, db.data[item])
}
func (db *database) create(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db.data[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400

		fmt.Fprintf(w, "duplicate item: %q\n", item)
		return
	}

	if f64, err := strconv.ParseFloat(price, 32); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400

		fmt.Fprintf(w, "invalid price: %q\n", price)
	} else {
		db.data[item] = dollars(f64)

		fmt.Fprintf(w, "added %s with price %s\n", item, dollars(f64))
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db.data[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404

		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	if f64, err := strconv.ParseFloat(price, 32); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400

		fmt.Fprintf(w, "invalid price: %q\n", price)
	} else {
		db.data[item] = dollars(f64)

		fmt.Fprintf(w, "new price %s for %s\n", dollars(f64), item)
	}
}

func (db *database) drop(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	item := req.URL.Query().Get("item")

	if _, ok := db.data[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404

		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db.data, item)

	fmt.Fprintf(w, "dropped %s\n", item)
}

func main() {
	db := database{
		data: map[string]dollars{
			"shoes": 50,
			"socks": 5,
		},
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/fetch", db.fetch)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/drop", db.drop)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
