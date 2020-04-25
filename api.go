package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	r.ParseForm()
	firstname := r.Form.Get("firstname")
	r.ParseForm()
	lastname := r.Form.Get("lastname")
	fmt.Println(id + " " + firstname + " " + lastname)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
