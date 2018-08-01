package main

import (
	"net/http"

	"./page"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Router)

	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("text/x-javascript", js.Minify)
	m.AddFunc("application/javascript", js.Minify)
	m.AddFunc("application/x-javascript", js.Minify)
	mux.Handle("/js/", http.StripPrefix("/js/", m.Middleware(http.FileServer(http.Dir("js")))))

	http.ListenAndServe("127.0.0.1:8000", mux)
}

// Router func
func Router(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form.Get("e") == "e" {
		page.Essence(w, r)
	} else {
		page.Full(w, r)
	}
}
