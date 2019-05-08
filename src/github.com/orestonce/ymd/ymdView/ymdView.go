package ymdView

import (
	"io"
)

type HtmlRenderer interface {
	HtmlRender(w io.Writer)
}

type HtmlRendererList []HtmlRenderer

func (this *HtmlRendererList) Add(a HtmlRenderer) {
	*this = append(*this, a)
}

func (thisObj HtmlRendererList) HtmlRender(w io.Writer) {
	for _, item := range thisObj {
		item.HtmlRender(w)
	}
}
