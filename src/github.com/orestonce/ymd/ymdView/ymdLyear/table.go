package ymdLyear

import (
	"github.com/orestonce/ymd/ymdView/ymdXss"
	"io"
	"fmt"
	"github.com/orestonce/ymd/ymdView"
)

type Table struct {
	TitleRow    []string
	BodyRowList [][]ymdView.HtmlRenderer
}

func (obj Table) HtmlRender(buf io.Writer) {
	tplTable(buf, obj)
}

type AddOneRowRequest struct {
	PreList  []string
	PostList []ymdView.HtmlRenderer
}

func (this *Table) AddOneRow(req AddOneRowRequest) {
	row := []ymdView.HtmlRenderer{}
	for _, one := range req.PreList {
		row = append(row, String(one))
	}
	for _, one := range req.PostList {
		row = append(row, one)
	}
	this.BodyRowList = append(this.BodyRowList, row)
}

func (this *Table) AddOneRowString(ss ...string) {
	this.AddOneRow(AddOneRowRequest{
		PreList: ss,
	})
}

func tplTable(buf io.Writer, obj Table) {
	buf.Write([]byte(`
                  <table class="table table-hover table-striped">
                    <thead>
                      <tr>`))
	for _, one := range obj.TitleRow {
		buf.Write([]byte("<th>"))
		buf.Write([]byte(ymdXss.Urlv(one)))
		buf.Write([]byte("</th>"))
	}
	buf.Write([]byte(`</tr>
                    </thead>
                    <tbody>
`))
	for _, row := range obj.BodyRowList {
		fmt.Fprint(buf, `<tr>`, "\n")
		for _, col := range row {
			fmt.Fprint(buf, "<td>")
			col.HtmlRender(buf)
			fmt.Fprint(buf, "</td>\n")
		}
		fmt.Fprint(buf, "</tr>\n")
	}
	fmt.Fprint(buf, `</tbody>
                  </table>
	`)
}
