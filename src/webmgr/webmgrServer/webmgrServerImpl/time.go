package webmgrServerImpl

import (
	"time"
	"sync"
	"github.com/orestonce/ymd/ymdError"
)

var localOnce sync.Once
var local *time.Location

func getTimeStr(t time.Time) string {
	localOnce.Do(func() {
		var err error
		local, err = time.LoadLocation(`Asia/Chongqing`)
		ymdError.PanicIfError(err)
	})
	return t.In(local).Format("2006-01-02 15:04:05")
}

