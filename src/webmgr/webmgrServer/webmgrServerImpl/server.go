package webmgrServerImpl

import (
	"github.com/orestonce/ymd/ymdView/ymdLyear"
	"log"
	"github.com/orestonce/ymd/ymdGin"
)

func RunServer(pwd string) {
	ymdGin.SetMode(ymdGin.ReleaseMode)

	engine := ymdGin.Default()
	ymdLyear.RegisterResourceToGinEngine(engine)
	engine.Use(func(ctx *ymdGin.Context) {
		switch ctx.Request.URL.Path {
		case adminUrl(`LoginPage`), adminUrl(`LoginAction`), adminUrl(`WsConnect`):
			break
		default:
			if !IsLoginSession(ctx) {
				ctx.RedirectTemp(adminUrl(`LoginPage`))
				ctx.Abort()
				log.Println(`Not login`, ctx.Request.RequestURI)
			}
		}
	})
	ymdGin.RegisterObjToEngine(engine, funcPrefix, adminWeb{})
	engine.NoRoute(func(ctx *ymdGin.Context) {
		ctx.RedirectTemp(adminUrl(`LoginPage`))
	})
	InitLoginCookieValue(pwd)
	engine.Run(":7611")
}
