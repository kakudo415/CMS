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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			r.URL.Path = "/index"
		}
		v := page.Get(r.URL.Path, r.FormValue("r")).Min()

		w.Write(v)
	})

	min := minify.New()
	min.AddFunc("image/svg+xml", svg.Minify)
	min.AddFunc("application/javascript", js.Minify)
	min.AddFunc("application/x-javascript", js.Minify)
	http.Handle("/static/", http.StripPrefix("/static/", min.Middleware(http.FileServer(http.Dir(os.Getenv("CMS_ROOT")+"static")))))

	http.ListenAndServe("127.0.0.1:8000", nil)
}
