package page

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/russross/blackfriday"
)

var indexHTML []byte

var headerRegex = regexp.MustCompile(`# .+`)

// Full Page Handler
func Full(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write(indexHTML)
		return
	}
	view := indexHTML
	content, title := parseArticle(r.URL.Path)
	view = bytes.Replace(view, []byte("[TITLE]"), title, 1)
	view = bytes.Replace(view, []byte("[BODY]"), content, 1)
	if len(view) > 0 {
		w.Write(view)
	} else {
		w.WriteHeader(404)
	}
}

// Essence data
func Essence(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("<!DOCTYPE html><head><title>404</title></head><body><h1>404</h1></body>"))
		return
	}
	content, title := parseArticle(r.URL.Path)
	w.Write([]byte("<!DOCTYPE html><head><title>" + string(title) + "</title></head><body>" + string(content) + "</body>"))
}

func parseArticle(filename string) ([]byte, []byte) {
	file, err := ioutil.ReadFile("md/" + filename + ".md")
	if err != nil {
		return []byte{}, []byte("404")
	}
	title := bytes.TrimPrefix(headerRegex.Find(file), []byte("# "))
	content := blackfriday.MarkdownBasic(file)
	if len(title) > 0 {
		return content, title
	}
	return content, []byte(filename) // <title>が見つからなかった場合応急処置としてファイル名を使う
}

func init() {
	var err error
	indexHTML, err = ioutil.ReadFile("view/index.html")
	if err != nil {
		panic(err)
	}
}
