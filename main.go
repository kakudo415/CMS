package main

import (
	"net/http"
	"strings"

	"./page"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			r.URL.Path = "/index"
		}
		v := page.Get(r.URL.Path, r.FormValue("r")).Min()

		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			v = v.Gzip()
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write(v)
	})
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.ListenAndServe("127.0.0.1:8000", nil)
}
