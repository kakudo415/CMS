package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe("127.0.0.1:8000", mux)
}

func router(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form.Get("e") == "e" {
		time.Sleep(time.Second * 1)
		w.Write([]byte("<!DOCTYPE html><head><title>hogehoge</title></head><body><p>hogehoge</p></body>"))
	} else {
		b, _ := ioutil.ReadFile("view/index.html")
		w.Write(b)
	}
}
