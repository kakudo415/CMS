package main

import (
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"./page"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		accessLog(r)
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

func accessLog(r *http.Request) {
	log.Printf("%s => %s %s", r.RemoteAddr, r.Method, r.URL.Path)
}

func init() {
	u, e := user.Current()
	if e != nil {
		panic(e)
	}
	f, e := os.OpenFile(filepath.Clean(u.HomeDir+"/.access"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if e != nil {
		panic(e)
	}
	log.SetOutput(f)
}
