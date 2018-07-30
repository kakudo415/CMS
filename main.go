package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)
	http.ListenAndServe("127.0.0.1:8000", mux)
}

func router(w http.ResponseWriter, r *http.Request) {
	println("!")
	r.ParseForm()
	if r.Form.Get("e") == "e" {
		w.Write([]byte("必要な情報だけを返します"))
	} else {
		b, _ := ioutil.ReadFile("view/index.html")
		w.Write(b)
	}
}
