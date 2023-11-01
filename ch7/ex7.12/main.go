// Exercise 7.12: Change the handler for /list to print its output as an HTML table, not text. You may find the html/template package(§4.6) useful.
package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const (
	templ = `
<html>

<head>
<h1>Items</h1>
</head>

<body>
<table>
<tr style="text-align: left">
    <th>Item</th>
    <th>Price</th>
</tr>

{{range $item, $price := .}}
<tr>
    <td>{{$item}}</td>
    <td>{{$price}}</td>
</tr>
{{end}}

</table>
</body>

</html>
`
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.New("list").Parse(templ))
	if err := tmpl.Execute(w, db); err != nil {
		log.Println(err)
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
