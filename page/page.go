package page

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"

	"github.com/russross/blackfriday"
)

var indexHTML []byte

var headerRegex = regexp.MustCompile(`# .+`)

var min = minify.New()

// Full Page Handler
func Full(w http.ResponseWriter, r *http.Request) {
	view := indexHTML
	if r.URL.Path != "/" {
		title, content := parseArticle(r.URL.Path)
		view = bytes.Replace(view, []byte("[TITLE]"), title, 1)
		view = bytes.Replace(view, []byte("[BODY]"), content, 1)
	}
	min.Minify("text/html", w, bytes.NewReader(view))
}

// Essence data
func Essence(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write(render([]byte("404"), []byte("<h1>404</h1><h2>Not Found</h2>")))
		return
	}
	title, content := parseArticle(r.URL.Path)
	min.Minify("text/html", w, bytes.NewReader(render(title, content)))
}

func parseArticle(filename string) ([]byte, []byte) {
	file, err := ioutil.ReadFile("Content/" + filename + ".md")
	if err != nil {
		return []byte{}, []byte("404")
	}
	title := bytes.TrimPrefix(headerRegex.Find(file), []byte("# "))
	content := blackfriday.MarkdownBasic(file)
	if len(title) > 0 {
		return title, content
	}
	return []byte(filename), content // <title>が見つからなかった場合応急処置としてファイル名を使う
}

func render(title, body []byte) (html []byte) {
	html = []byte("<!DOCTYPE html><html><meta charset=\"UTF-8\"><head><title>")
	html = append(html, title...)
	html = append(html, []byte("</title></head><body>")...)
	html = append(html, body...)
	html = append(html, []byte("</body></html>")...)
	return html
}

func init() {
	var err error
	indexHTML, err = ioutil.ReadFile("view/index.html")
	if err != nil {
		panic(err)
	}

	min.AddFunc("text/html", html.Minify)
	min.AddFunc("text/css", css.Minify)
}
