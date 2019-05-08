package ymdOs

import (
	"testing"
)

func TestMustGetWd(t *testing.T) {
	MustGetWd()
}

func TestMustChdir(t *testing.T) {
	MustChdir("/")
}
