package ymdLyear

import (
	"io"
	"fmt"
	"github.com/orestonce/ymd/ymdUuid"
	"github.com/orestonce/ymd/ymdView/ymdXss"
)

type InputString struct {
	Name        string
	ShowName    string
	Value       string
	Placeholder string
	IsPassword  bool
	ReadOnly    bool
}

func (obj InputString) HtmlRender(w io.Writer) {
	Id := ymdUuid.MustNewUUID()
	if obj.ShowName == `` {
		obj.ShowName = obj.Name
	}
	fmt.Fprint(w, `<div class="form-group">`, "\n")
	fmt.Fprint(w, `    <label class="col-md-3 control-label" for="`, ymdXss.Urlv(Id), `">`, ymdXss.Urlv(obj.ShowName), `</label>`, "\n")
	fmt.Fprint(w, `    <div class="col-md-7"><input `, )
	if obj.IsPassword {
		fmt.Fprint(w, `type="password" `)
	}
	if obj.Value != `` {
		fmt.Fprint(w, `value="`, ymdXss.Urlv(obj.Value), `" `)
	}
	if obj.ReadOnly {
		fmt.Fprint(w, `readonly="readonly" `)
	}
	fmt.Fprint(w, `class="form-control" id="`, ymdXss.Urlv(Id), `" name="`, ymdXss.Urlv(obj.Name), `" placeholder="`, ymdXss.Urlv(obj.Placeholder), `"></div></div>`, "\n")
}
