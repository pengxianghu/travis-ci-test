package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io"
	"encoding/json"
)

type result struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data string `json:"data"`
}

type middleWareHandler struct {
	r *httprouter.Router
}

func registerHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/test", testHandler)

	return router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	m.r.ServeHTTP(w, r)
}

func main() {
	r := registerHandler()
	mh := NewMiddleWareHandler(r)

	http.ListenAndServe(":6000", mh)
}


func testHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var res result
	res.Code = 200
	res.Msg = "success"
	res.Data = "get handler"

	resp, _ := json.Marshal(res)
	io.WriteString(w, string(resp))
}
