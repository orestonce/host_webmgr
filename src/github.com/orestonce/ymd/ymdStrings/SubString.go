package ymdStrings

import "strings"

func StringAfterLastSubString(total string, sub string) string {
	idx := strings.LastIndex(total, sub)
	if idx < 0 || idx+len(sub) >= len(total) {
		return ``
	}
	return total[idx+len(sub):]
}
