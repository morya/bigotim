package main

type MsgHead struct {
	Length  uint32
	Cmd     uint32
	Status  uint16
	Version uint16 // 为扩展保留
	Seq     uint32
}

type MsgAck struct {
	MsgHead
}

type MsgLogin struct {
	Account string
	Hash    string
}

type MsgHeartBeat struct {
	MsgHead
}
