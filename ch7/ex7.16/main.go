// Exercise 7.16: Write a web-based calculator program.
package main

import (
	"fmt"
	"log"
	"net/http"

	"gopl/ch7/examples/eval"
)

func main() {
	http.HandleFunc("/", calculate)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func calculate(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("eval")
	expr, err := eval.Parse(s)
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	env := make(eval.Env)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
	}
	fmt.Fprintf(w, "%v\n", expr.Eval(env))
}
