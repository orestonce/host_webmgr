package ymdLyear

import (
	"github.com/orestonce/ymd/ymdView"
	"io"
	"fmt"
	"github.com/orestonce/ymd/ymdView/ymdXss"
)

type Card struct {
	Title string
	Body  ymdView.HtmlRenderer
}

func (obj Card) HtmlRender(w io.Writer) {
	fmt.Fprint(w, `<div class="card">`, "\n")
	fmt.Fprint(w, `<div class="card-header"><h4>`, ymdXss.Urlv(obj.Title), `</h4></div>`, "\n")
	fmt.Fprint(w, `<div class="card-body">`, "\n")
	obj.Body.HtmlRender(w)
	fmt.Fprint(w, "</div>\n", "</div>\n", "</div>\n")
}

type Pre string

func (obj Pre) HtmlRender(w io.Writer) {
	fmt.Fprint(w, `<pre>`, ymdXss.Urlv(string(obj)), `</pre>`)
}
