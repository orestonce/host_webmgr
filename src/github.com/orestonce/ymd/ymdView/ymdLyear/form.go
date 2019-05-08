package ymdLyear

import (
	"io"
	"fmt"
	"github.com/orestonce/ymd/ymdView/ymdXss"
	"github.com/orestonce/ymd/ymdView"
)

type Form struct {
	TopTitle  string
	ActionUrl string
	IsPost    bool
	InputList []ymdView.HtmlRenderer
}

func (obj Form) HtmlRender(w io.Writer) {
	method := `GET`
	if obj.IsPost {
		method = `POST`
	}
	fmt.Fprint(w, `<div class="card">`, "\n")
	fmt.Fprint(w, `<div class="card-header"><h4>`, ymdXss.Urlv(obj.TopTitle), `</h4></div>`, "\n")
	fmt.Fprint(w, `<div class="card-body">
                			<form class="form-horizontal" action="`, obj.ActionUrl, `" method="`, method, `">`)
	for _, one := range obj.InputList {
		one.HtmlRender(w)
	}
	fmt.Fprint(w, `
				  <div class="form-group text-center">
                    <button class="btn btn-primary btn-w-xl" type="submit">提交</button>
                  </div>
</div>
                </form>
				</div>
</div>
`)
}
