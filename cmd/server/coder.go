package main

import (
	"bytes"
	"encoding/hex"

	log "github.com/CodisLabs/codis/pkg/utils/log"
)

type Coder struct {
}

func NewCoder() *Coder {
	return &Coder{}
}

func (this *Coder) Marshal(buffer []byte, msg interface{}) {
	switch realmsg := msg.(type) {
	case MsgHeartBeat:
		this.marshalHeartBeat(buffer, realmsg)

	case MsgLogin:
		this.marshalLogin(buffer, realmsg)

	case MsgAck:
		this.marshalMsgAck(buffer, realmsg)
	}
}

func (c *Coder) marshalHeartBeat(buffer []byte, msg MsgHeartBeat) {
	log.Println("marshalHeartBeat called")

}

func (c *Coder) marshalLogin(buffer []byte, msg MsgLogin) {
	log.Println("marshalLogin called")
	log.Println("account:", msg.Account)
	log.Println("hash:", msg.Hash)
	var b bytes.Buffer
	dumper := hex.Dumper(&b)
	//writer := bufio.NewWriter(dumper)
	dumper.Write([]byte{'\013', '\013', '\013', '\010', '\022'})
	log.Println("acc hex:", string(b.Bytes()))

}

func (c *Coder) marshalMsgAck(buffer []byte, msg MsgAck) {
	log.Println("marshalMsgAck called")
}
