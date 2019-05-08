package ymdLyear

import (
	"io"
	"fmt"
	"github.com/orestonce/ymd/ymdView/ymdXss"
)

type Button struct {
	Href                string
	ShowContent         string
	IsDanger            bool
	IsAjaxSendAndReload bool
	ConfirmText         string
}

func (obj Button) HtmlRender(w io.Writer) {
	more := `btn-success`
	if obj.IsDanger {
		more = `btn-danger`
	}
	fmt.Fprint(w, `<a class="btn `, more, ` btn-w-md" type="button" `)
	if obj.IsAjaxSendAndReload {
		fmt.Fprint(w, ` onclick=ymd_send_and_reload("`, obj.Href, `") `)
	} else {
		fmt.Fprint(w, ` href="`, obj.Href, `" `)
	}
	fmt.Fprint(w, ">", ymdXss.Urlv(obj.ShowContent), `</a>`)
}
