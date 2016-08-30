package main

import (
	log "github.com/CodisLabs/codis/pkg/utils/log"

	"flag"
)

var (
	remoteHost  = flag.String("host", "127.0.0.1:8899", "msgserver address")
	clientCount = flag.Uint("conn", 1, "connection made to server")
	sleepDebug  = flag.Bool("sleep", true, "debug sleep, will sleep in loop")
)

func main() {
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	linkmng := newLinkMng(*remoteHost)
	linkmng.Main()
}
