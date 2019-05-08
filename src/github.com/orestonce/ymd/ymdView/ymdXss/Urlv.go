package ymdXss

import "html"

func Urlv(s string) string {
	return html.EscapeString(s)
}
