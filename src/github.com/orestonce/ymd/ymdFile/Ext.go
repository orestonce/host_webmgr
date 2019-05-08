package ymdFile

import (
	"io/ioutil"
)

func MustWriteFile(FileName string, Content []byte) {
	err := ioutil.WriteFile(FileName, Content, 0600)
	if err != nil {
		panic(`wpw4bxc2 ` + err.Error())
	}
}
