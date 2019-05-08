package ymdView

import (
	"bytes"
	"testing"
)

func TestHtmlRendererList_Add(t *testing.T) {
	var list HtmlRendererList
	list.Add(list)
}

func TestHtmlRendererList_HtmlRender(t *testing.T) {
	var list HtmlRendererList
	list.HtmlRender(bytes.NewBuffer(nil))
}
