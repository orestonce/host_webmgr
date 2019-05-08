package ymdLyear_Build

import (
	"github.com/orestonce/ymd/ymdBindata"
	"strings"
)

func YmdBuild() {
	ymdBindata.MustBuildResource(ymdBindata.MustBuildResourceRequest{
		Source: `src/github.com/orestonce/ymd/ymdView/ymdLyear/ymdLyear_Build/front`,
		Output: `src/github.com/orestonce/ymd/ymdView/ymdLyear/zzzig_BuildFront.go`,
		SkipFileCb: func(filename string) bool {
			return strings.HasSuffix(filename, `.html`)
		},
	})
}
