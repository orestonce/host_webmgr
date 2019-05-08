package main

import (
	"log"
	"flag"
	"webmgr/webmgrServer/webmgrServerImpl"
)

func main() {
	var pwd string
	flag.StringVar(&pwd, `p`, ``, `password`)
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	webmgrServerImpl.RunServer(pwd)
}
