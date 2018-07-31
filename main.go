package main

import (
	"net/http"
	"time"

	"./page"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Router)
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.ListenAndServe("127.0.0.1:8000", mux)
}

// Router func
func Router(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form.Get("e") == "e" {
		time.Sleep(time.Second * 1)
		page.Essence(w, r)
	} else {
		page.Full(w, r)
	}
}
