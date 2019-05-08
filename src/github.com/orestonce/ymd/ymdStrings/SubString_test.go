package ymdStrings

import (
	"testing"
	"github.com/orestonce/ymd/ymdAssert"
)

func TestStringAfterLastSubString(t *testing.T) {
	after := StringAfterLastSubString(`1.2.3.456`, `.`)
	ymdAssert.True(after == `456`, after)
	after = StringAfterLastSubString(`1.2.3.45`, `go`)
	ymdAssert.True(after == ``)
	after = StringAfterLastSubString(`1.2.3.45`, `45`)
	ymdAssert.True(after == ``, after)
}
