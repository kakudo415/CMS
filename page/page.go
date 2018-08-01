package page

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"

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
	var title, content []byte
	if r.URL.Path == "/" {
		title, content = []byte("Kakudo's Blog"), list("Content")
	} else {
		title, content = parseArticle(r.URL.Path)
	}
	view = bytes.Replace(view, []byte("[TITLE]"), title, 1)
	view = bytes.Replace(view, []byte("[BODY]"), content, 1)
	min.Minify("text/html", w, bytes.NewReader(view))
}

// Essence data
func Essence(w http.ResponseWriter, r *http.Request) {
	var title, content []byte
	if r.URL.Path == "/" {
		title, content = []byte("Kakudo's Blog"), list("Content")
	} else {
		title, content = parseArticle(r.URL.Path)
	}
	min.Minify("text/html", w, bytes.NewReader(render(title, content)))
}

func parseArticle(filename string) ([]byte, []byte) {
	file, err := ioutil.ReadFile(filepath.Clean("Content/" + filename + ".md"))
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

func list(d string) (l []byte) {
	files, _ := ioutil.ReadDir(d)
	for _, file := range files {
		p := filepath.Join(d, file.Name())
		p = trimExt(strings.TrimPrefix(p, "Content\\"))
		if file.IsDir() {
			l = append(l, list(p)...)
		} else {
			t, c := parseArticle(string(p))
			var wrapUp string
			if utf8.RuneCount(c) > 100 {
				wrapUp = string([]rune(string(c))[:100])
			} else {
				wrapUp = string([]rune(string(c))[:utf8.RuneCount(c)])
			}
			wrapUp = strings.TrimPrefix(removeTag(wrapUp), string(t))
			wrapUp = string([]rune(wrapUp)[utf8.RuneCount(t):])
			l = append(l, []byte(`<a href="/`+trimExt(file.Name())+`"><h1>`+string(t)+`</h1><span>`+wrapUp+`</span></a>`)...)
		}
	}
	return l
}

func removeTag(str string) string {
	rep := regexp.MustCompile(`<("[^"]*"|'[^']*'|[^'">])*>`)
	str = rep.ReplaceAllString(str, "")
	return str
}

func trimExt(p string) string {
	return strings.TrimSuffix(p, filepath.Ext(p))
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
