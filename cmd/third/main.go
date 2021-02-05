package main

import (
	"io/ioutil"
	"net/http"
)

func handler (res http.ResponseWriter, req *http.Request) {
	if http.MethodPost == req.Method {
		reqBody, _ := ioutil.ReadAll(req.Body)
		setHeaders(res, req)
		res.Write(reqBody)
	} else {
		res.Write([]byte("Please, choose POST method."))
	}
}

func setHeaders(res http.ResponseWriter, req *http.Request) {
	reqHeaders := req.Header
	for key, values := range reqHeaders {
		for _, value := range values {
			res.Header().Set(key, value)
		}
	}
}

func main() {
	http.HandleFunc("/post", handler)
	http.ListenAndServe(":8080", nil)
}