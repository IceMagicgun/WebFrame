package main

import (
	"WebFrame/core"
	"net/http"
)

func main() {
	myRouter := core.NewRouter()
	http.ListenAndServe("0.0.0.0:8068", myRouter)
}
