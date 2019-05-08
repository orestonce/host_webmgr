package ymdLyear

import (
	"io"
	"fmt"
	"strings"
)

type Blank struct {
	Count int
}

func (obj Blank) HtmlRender(w io.Writer) {
	fmt.Fprint(w, strings.Repeat(`&nbsp;`, obj.Count))
}
