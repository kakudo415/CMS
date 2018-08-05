package main

import (
	"net/http"
	"os"

	"./page"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/svg"
)

func main() {
	go Resource()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			r.URL.Path = "/index"
		}
		v := page.Get(r.URL.Path, r.FormValue("r")).Min()

		w.Write(v)
	})

	http.ListenAndServe("127.0.0.1:8000", nil)
}

// Resource Server
func Resource() {
	min := minify.New()
	min.AddFunc("image/svg+xml", svg.Minify)
	min.AddFunc("application/javascript", js.Minify)
	min.AddFunc("application/x-javascript", js.Minify)

	mux := http.NewServeMux()
	mux.Handle("/", min.Middleware(http.FileServer(http.Dir(dir()+"static"))))
	http.ListenAndServe("127.0.0.1:10000", mux)
}

func dir() string {
	if d := os.Getenv("CMS_ROOT"); len(d) > 0 {
		return d
	}
	panic("ROOT DIRECTORY NOT FOUND")
}
