package ymdOs

import (
	"os"
)

func MustGetWd() string {
	wds, err := os.Getwd()
	if err != nil {
		panic("jsvji875 " + err.Error())
	}
	return wds
}

func MustChdir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		panic("meop9bd1 " + err.Error())
	}
}
