package main

import (
	log "github.com/CodisLabs/codis/pkg/utils/log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c := NewCoder()
	buff := make([]byte, 256)
	msg := MsgLogin{
		MsgHead: MsgHead{
			Length:  0,
			Cmd:     0,
			Status:  15,
			Version: 0,
			Seq:     1,
		},
		Account: "morya",
		Hash:    "addfadfacefadf12358adf",
	}
	c.Marshal(buff, msg)

	log.Println("Bye")
}
