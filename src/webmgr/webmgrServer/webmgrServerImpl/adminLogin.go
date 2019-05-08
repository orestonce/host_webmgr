package webmgrServerImpl

import (
	"github.com/orestonce/ymd/ymdGin"
	"github.com/orestonce/ymd/ymdView/ymdLyear"
	"crypto/rand"
	"github.com/orestonce/ymd/ymdError"
	"encoding/hex"
	"log"
	"crypto/md5"
)

const loginFlag = `loginFlag`

var gLoginCookieValue = ``

func InitLoginCookieValue(pwd string) {
	if pwd == `` {
		data := make([]byte, 10)
		_, err := rand.Read(data)
		ymdError.PanicIfError(err)
		pwd = hex.EncodeToString(data)
	}
	gLoginCookieValue = getMd5Hex(`admin:` + pwd)
	log.Println(`Init login cookie ok, username/password : admin/` + pwd)
}

func IsLoginSession(ctx *ymdGin.Context) bool {
	return ctx.Cookie(loginFlag) == gLoginCookieValue
}

func DelLoginSession(ctx *ymdGin.Context) {
	ctx.SetCookieSimple(loginFlag, `nil`)
}

func TryLogin(ctx *ymdGin.Context, req ymdLyear.LoginActionRequest) (isLogin bool) {
	if !IsLoginSession(ctx) {
		if gLoginCookieValue == getMd5Hex(req.Username+":"+req.Password) {
			ctx.SetCookieSimple(loginFlag, gLoginCookieValue)
			isLogin = true
		}
	} else {
		isLogin = true
	}
	return
}

func getMd5Hex(origin string) string {
	m := md5.New()
	m.Write([]byte(origin))
	r := m.Sum(nil)
	return hex.EncodeToString(r)
}
