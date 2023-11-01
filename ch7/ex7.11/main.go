// Exercise 7.11: Add aditional handlers so that clients can create, read, update, and delete database entries. For example, a request of the form /update?item=socks&price=6 will update the price of an item in the inventory and report an error if the item does not exist or if the price is invalid. (Warning: this change introduces concurrent variable updates.)
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/read", db.read)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
		if err != nil {
			log.Fatal(err)
		}
		db[item] = dollars(price)
		fmt.Fprintf(w, "Added %s: %s\n", item, db[item])
		return
	}
	w.WriteHeader(http.StatusNotFound) // 404
	fmt.Fprintf(w, "Item already exists\n")
	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	nprice, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		log.Fatal(err)
	}
	db[item] = dollars(nprice)
	fmt.Fprintf(w, "%s: %s -> %s\n", item, price, db[item])
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "%s has been deleted\n", item)
}

