package main

import (
	"fmt"
	"net/http"
)

type Counter int

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	*c++
	fmt.Fprintf(w, "counter = %d\n", *c)
}

func main() {
	ctr := new(Counter)
	http.Handle("/counter", ctr)
	http.ListenAndServe(":8080", nil)
}
