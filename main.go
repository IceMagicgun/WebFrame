package main

import (
	"WebFrame/core"
	"net/http"
)

func main() {
	myRouter := core.NewRouter()
	http.ListenAndServe("localhost:8808", myRouter)
}
