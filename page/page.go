package page

import (
	"bytes"
	"io/ioutil"
	"regexp"

	"github.com/russross/blackfriday"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

// Response type
type Response []byte

var hRegex = regexp.MustCompile(`# .+`)
var min = minify.New()

// Get Completed HTML
func Get(p, t string) Response {
	var v []byte
	var e error
	if t == "i" {
		v = []byte(`<!DOCTYPE html><html lang="ja"><head><meta charset="UTF-8"><title>[TITLE]</title></head><body>[CONTENT]</body></html>`)
	} else {
		v, e = ioutil.ReadFile("view/index.html")
	}

	c, e := ioutil.ReadFile("Content" + p + ".md")
	if e != nil {
		return []byte(`<!doctype html><html><body><h1>404</h1></body></html>`)
	}

	title, content := bytes.TrimPrefix(hRegex.Find(c), []byte("# ")), blackfriday.MarkdownBasic(c)
	if p == "/index" {
		title = []byte(`Kakudo's Blog`)
	} else {
		content = []byte(`<div id="article">` + string(content) + `</div>`)
	}
	v = bytes.Replace(v, []byte("[TITLE]"), title, 1)
	v = bytes.Replace(v, []byte("[CONTENT]"), content, 1)

	return v
}

// Min func for HTML Minify
func (r Response) Min() Response {
	src := bytes.NewBuffer(r)
	dst := bytes.NewBuffer([]byte{})
	min.Minify("text/html", dst, src)
	return dst.Bytes()
}

func init() {
	min.AddFunc("text/html", html.Minify)
	min.AddFunc("text/css", css.Minify)
	min.AddFunc("application/javascript", js.Minify)
}
