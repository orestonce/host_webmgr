package main

import (
	"github.com/orestonce/ymd/ymdIpToRegion/ymdIpToRegion_Build"
	"github.com/orestonce/ymd/ymdView/ymdLyear/ymdLyear_Build"
)

func main() {
	ymdIpToRegion_Build.YmdBuild()
	ymdLyear_Build.YmdBuild()
}