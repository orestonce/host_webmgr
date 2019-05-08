package ymdLyear

import (
	"io"
	"fmt"
	"github.com/orestonce/ymd/ymdView/ymdXss"
)

type String string

func (obj String) HtmlRender(w io.Writer) {
	fmt.Fprint(w, ymdXss.Urlv(string(obj)))
}
