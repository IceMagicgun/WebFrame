package core

import (
	"bytes"
	"encoding/json"
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
	ans["POST"]["/GetData"] = _GetData
	ans["GET"]["/GetData"] = _GetData
	return &MyRouter{ans}
}

func (this *MyRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var f http.HandlerFunc
	println(req.Method, "  ", req.URL.Path)
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

var (
	HasGetData = false
	Data       = []byte{}
)

func GetData(res http.ResponseWriter, req *http.Request) {
	if !HasGetData {
		reqData := map[string]interface{}{}
		reqData["ServerId"] = "5"
		reqBody, _ := json.Marshal(reqData)
		res, _ := http.Post("http://47.97.90.13:8808/Admin/GetTaskSideData", "application/json", bytes.NewReader(reqBody))
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		//println(string(body))
		data := map[string]interface{}{}
		json.Unmarshal(body, &data)
		Data, _ = json.Marshal(data["Data"])
		//println(string(Data))
		HasGetData = true
	}
	res.Write(Data)
}

func _GetData(res http.ResponseWriter, req *http.Request) {
	content, _ := ioutil.ReadFile("./other/数据.txt")
	data := map[string]interface{}{}
	json.Unmarshal(content, &data)
	Data, _ = json.Marshal(data["Data"])
	//println(string(Data))
	HasGetData = true
	res.Write(Data)
}
