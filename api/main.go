package main

import (
	"net/http"
)

func main() {
	r := newRouter()

	http.ListenAndServe(":3000", r)
}
