package main

import (
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"./page"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		accessLog(r)
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
