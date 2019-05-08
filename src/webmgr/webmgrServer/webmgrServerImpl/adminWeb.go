package webmgrServerImpl

import (
	"github.com/orestonce/ymd/ymdGin"
	"github.com/orestonce/ymd/ymdView/ymdLyear"
	"github.com/orestonce/ymd/ymdView"
	"github.com/orestonce/ymd/ymdView/ymdXss"
	"bytes"
	"io"
)

const funcPrefix = `/admin/`

type adminWeb struct{}

func (adminWeb) LoginPage(ctx *ymdGin.Context) {
	if IsLoginSession(ctx) {
		ctx.RedirectTemp(adminUrl(`WsClientListPage`))
		return
	}
	ymdLyear.TplLoginPage(ctx.Writer, ymdLyear.TplLoginPageRequest{
		ActionUrl: adminUrl(`LoginAction`),
		ErrMsg:    ctx.InStr(`err`),
	})
}

func (adminWeb) LoginAction(ctx *ymdGin.Context) {
	req := ymdLyear.LoginActionGetRequest(ctx)
	if !TryLogin(ctx, req) {
		ctx.RedirectTemp(adminUrlArgs(`LoginPage`, map[string]string{
			`err`: `密码错误`,
		}))
		return
	}
	ctx.RedirectTemp(adminUrl(`WsClientListPage`))
}

func (adminWeb) LogoutAction(ctx *ymdGin.Context) {
	DelLoginSession(ctx)
	ctx.RedirectTemp(adminUrl(`LoginPage`))
}

func adminUrl(name string) string {
	return funcPrefix + ymdXss.Urlv(name)
}

func adminUrlArgs(name string, arg map[string]string) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(adminUrl(name))
	isFirst := true
	for key, value := range arg {
		if isFirst {
			buf.WriteString(`?`)
			isFirst = false
		} else {
			buf.WriteString(`&`)
		}
		buf.WriteString(ymdXss.Urlv(key))
		buf.WriteString(`=`)
		buf.WriteString(ymdXss.Urlv(value))
	}
	return buf.String()
}

func adminUi(w io.Writer, name string, main ymdView.HtmlRenderer) {
	ymdLyear.WriteHtmlPage(w, ymdLyear.WarpHtmlPage{
		Username:  `Admin`,
		LogoutUrl: adminUrl(`LogoutAction`),
		LeftSide: []ymdLyear.SideNode{
			{
				UrlPath:  adminUrl(`WsClientListPage`),
				ShowName: `ws客户端列表`,
			},
		},
	}, name, main)
}
