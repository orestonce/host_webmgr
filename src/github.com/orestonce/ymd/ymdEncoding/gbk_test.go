package ymdEncoding

import (
	"testing"
	"github.com/orestonce/ymd/ymdError"
	"github.com/orestonce/ymd/ymdAssert"
)

func TestUtf8ToGbk(t *testing.T) {
	const oldStr = `张三`
	gbkContent, err := Utf8ToGbk([]byte(oldStr))
	ymdError.PanicIfError(err)

	ymdAssert.True(oldStr != string(gbkContent))

	utf8Content, err := GbkToUtf8(gbkContent)
	ymdError.PanicIfError(err)
	ymdAssert.True(string(utf8Content) == oldStr)
}
