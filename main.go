package main

import (
	. "./controller"
	. "./router"
	"net/http"
)

func main() {
	Ctrl()

	r := GetRoute()

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
