package main

import (
	"net/http"

	"./page"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			r.URL.Path = "/index"
		}
		v := page.Get(r.URL.Path, r.FormValue("r")).Min()

		w.Write(v)
	})

	min := minify.New()
	min.AddFunc("application/javascript", js.Minify)
	min.AddFunc("application/x-javascript", js.Minify)
	http.Handle("/js/", http.StripPrefix("/js/", min.Middleware(http.FileServer(http.Dir("js")))))

	http.ListenAndServe("127.0.0.1:8000", nil)
}
