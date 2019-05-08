package ymdEncoding

import "golang.org/x/text/encoding/simplifiedchinese"

func GbkToUtf8(gbkContent []byte) (utf8Content []byte, err error) {
	return simplifiedchinese.GBK.NewDecoder().Bytes(gbkContent)
}

func Utf8ToGbk(utf8Content []byte) (gbkContent []byte, err error) {
	return simplifiedchinese.GBK.NewEncoder().Bytes(utf8Content)
}