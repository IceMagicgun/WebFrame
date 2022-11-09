package core

import (
	"io/ioutil"
	"net/http"
)

type MyRouter struct {
	Mapping map[string]map[string]http.HandlerFunc
}

func NewRouter() *MyRouter {
	return &MyRouter{make(map[string]map[string]http.HandlerFunc)}
}

func (this *MyRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var f http.HandlerFunc
	if handlerFuncs, has := this.Mapping[req.Method]; has {
		f = handlerFuncs[req.URL.Path]
	}
	if f != nil {
		f(res, req)
		return
	}
	content, err := ioutil.ReadFile("./html" + req.URL.Path + ".html")
	if err == nil {
		res.Write(content)
		return
	}
	content, err = ioutil.ReadFile("." + req.URL.Path)
	if err == nil {
		res.Write(content)
		return
	}
	content, _ = ioutil.ReadFile("./html/404.html")
	res.Write(content)
}
