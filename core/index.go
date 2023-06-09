package core

import (
	debugM "WebFrame/core/debug"
	"WebFrame/module/open_ai"
	"io/ioutil"
	"net/http"
)

type MyRouter struct {
	Mapping map[string]map[string]http.HandlerFunc
}

func NewRouter() *MyRouter {
	ans := map[string]map[string]http.HandlerFunc{}
	ans["POST"] = map[string]http.HandlerFunc{}
	ans["GET"] = map[string]http.HandlerFunc{}
	ans["POST"]["/GetData"] = open_ai.AskForever
	ans["GET"]["/GetData"] = open_ai.AskForever
	ans["POST"]["/GetHistory"] = open_ai.GetHistory
	ans["GET"]["/GetHistory"] = open_ai.GetHistory
	return &MyRouter{ans}
}

func (myRouter *MyRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var f http.HandlerFunc
	debugM.Log(req.Method + "  " + req.URL.Path)
	if handlerFuncs, has := myRouter.Mapping[req.Method]; has {
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
