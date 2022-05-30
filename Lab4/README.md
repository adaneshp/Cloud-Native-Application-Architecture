# Lab 4: Go Webserver
- This code was developed in February, 2022.

## Go web server
Building a Go web server with the net/http package (from The Go Programming Language book)

Here’s the basic outline - 
```
package http

type handler interface {
	ServerHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address, string, h Handler) error
```
The ListenAndServer function requires a server address such as “localhost:8000”, and an instance of the Handler interface to which all requests should be dispatched. It runs forever, or until the server fails.

To enable a user defined data type to be able to serve http requests, we need to attach a ServeHTTP method to it. 
For example, consider a map of (items:price) that we want to serve http requests from
```
package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
```

Install curl if needed
```
sudo apt update
sudo apt install curl
```
Run the webserver

```
go run webserver1.go
In a second terminal, use a curl client to access the server
curl “http://localhost:8000”
```

Note: If you are copy-pasting, the quotes may not copy correctly.

To incorporate multiple URL endpoints, we could use a switch statement
```
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}
```
Test with curl

```
curl “http://localhost:8000/list”
curl “http://localhost:8000/price?item=socks”
```

Alternatively, we could use a request multiplexer to simplify the association between URLs and handlers. A ServeMux aggregates a collection of http.Handlers into a single http.Handler
```

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

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
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
```

Note:
db.list and db.price do not have a ServerHTTP method, and hence cannot be passed directly to mux.Handle. So an adapter function is needed to add the ServerHTTP method (decorator pattern). The net/http package defines the HandlerFunc type for this purpose
```
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServerHTTP(w ResponseWriter, r *Request {
	f(w, r)
}
```

HandlerFunc is a function type that has methods that satisfies the http.Handler interface. Its behavior is to call the underlying function.

ServerMux has a convenience function HandleFunc to simplify the code for handler registration
```
mux.HandleFunc(“/list”, db.list)
mux.HandleFunc(“/price”, db.price)
```
Important:
The web server invoked each handler in a new goroutine, so take precautions to handle concurrent operations on shared data structures using locks and channels.

## Extra

 Added additional handlers so that clients can create, read, update, and delete (CRUD) database entries. For example, a request of the form /update?item=socks&price=6 will update the price of an item in the inventory and report an error if the item does not exist or if the price is invalid.

