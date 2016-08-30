package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"

	log "github.com/CodisLabs/codis/pkg/utils/log"
)

type Coder struct {
}

func NewCoder() *Coder {
	return &Coder{}
}

func (this *Coder) Marshal(msg interface{}) ([]byte, error) {
	var data []byte
	var err error

	switch realmsg := msg.(type) {
	case MsgHeartBeat:
		data, err = this.marshalHeartBeat(realmsg)

	case MsgLogin:
		data, err = this.marshalLogin(realmsg)

	case MsgAck:
		data, err = this.marshalMsgAck(realmsg)
	}

	if *sleepDebug {
		log.Printf("msg\n%s", hex.Dump(data))
	}

	return data, err
}

func (c *Coder) marshalHeartBeat(msg MsgHeartBeat) ([]byte, error) {
	log.Println("marshalHeartBeat called")

	return []byte{}, nil

}

func (c *Coder) marshalLogin(msg MsgLogin) ([]byte, error) {
	var msghead MsgHead
	var b bytes.Buffer

	//	log.Debug("account:", msg.Account)
	//	log.Debug("hash:", msg.Hash)

	logjson, _ := json.Marshal(msg)
	// out.Write(logjson)
	msghead.Length = uint32(len(logjson))
	msghead.Cmd = 3
	msghead.Status = 4
	msghead.Version = 5
	msghead.Seq = 6
	binary.Write(&b, binary.BigEndian, msghead)
	b.Write(logjson)

	return b.Bytes(), nil
}

func (c *Coder) marshalMsgAck(msg MsgAck) ([]byte, error) {
	log.Println("marshalMsgAck called")
	return []byte{}, nil
}
