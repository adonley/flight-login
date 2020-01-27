package main

import (
	"net/http"
)

func main() {
	router := NewRouter()
	_ = http.ListenAndServe(":3000", router)
}
