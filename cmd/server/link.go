package main

import (
	"bufio"
	"net"
)

type Link struct {
	netConn net.Conn

	reader *bufio.Reader
	writer *bufio.Writer
}

func NewLink(conn net.Conn) *Link {
	return &Link{
		netConn: conn,
		reader:  bufio.NewReader(conn),
		writer:  bufio.NewWriter(conn),
	}
}

func (l *Link) Run() {

}
