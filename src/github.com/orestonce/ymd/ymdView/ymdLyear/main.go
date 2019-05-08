package ymdLyear

import (
	"net/http"
	"github.com/orestonce/ymd/ymdGin"
	"strings"
)

func RegisterResourceToGinEngine(engine *ymdGin.Engine) {
	engine.GET(`/lyear/*respath`, func(ctx *ymdGin.Context) {
		respath := ctx.Param(`respath`)
		data, exists := GetBinData(respath)
		if !exists {
			ctx.AbortWithStatus(404)
			return
		}
		http.ServeContent(ctx.Writer, ctx.Request, respath, data.ModTime, strings.NewReader(data.Content))
	})
}
