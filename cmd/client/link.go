// client.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/CodisLabs/codis/pkg/utils/log"
)

type Link struct {
	exitChan chan struct{}
	mng      *LinkMng
	runid    int
	NetConn  net.Conn
	Reader   *bufio.Reader
	Writer   *bufio.Writer
	coder    Coder
}

///**
//* create link from net.Conn, for Server interface
// */
//func NewLink(id int, conn net.Conn) (*Link, error) {
//	return &Link{
//		runid:   id,
//		NetConn: conn,
//		Reader:  bufio.NewReader(conn),
//		Writer:  bufio.NewWriter(conn),
//	}, nil
//}

// create link from ip address, for client interface
func NewLinkDial(id int, addr string, l *LinkMng) (*Link, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		return nil, err
	}

	log.Print("dial connection ok")

	return &Link{
		exitChan: make(chan struct{}, 100),
		mng:      l,
		runid:    id,
		NetConn:  conn,
		Reader:   bufio.NewReader(conn),
		Writer:   bufio.NewWriter(conn),
	}, nil
}

func (c *Link) generateContent() string {
	return fmt.Sprintf("%v", time.Now().Unix())
}

func (c *Link) Run() {

	defer func() {
		if r := recover(); r != nil {
			log.Errorf("recover from err, %v", r)
		}
	}()

	go c.readLoop()

	writer := c.Writer
	var count int = 0

	for {

		msg := MsgLogin{
			Account: "morya",
			Hash:    c.generateContent(),
		}

		data, err := c.coder.Marshal(msg)
		if err != nil {
			// terminate this connection
			log.InfoError(err, "marshal msg failed.")
			return
		}

		c.NetConn.SetWriteDeadline(time.Now().Add(time.Millisecond * 300))
		writer.Write(data)
		writer.Flush()

		count += 1
		if count%1000 == 0 {
			// log.Println("sent login cmd, count ", count)
			c.mng.addStats(c.runid, count)
		}

		if *sleepDebug {
			time.Sleep(time.Second * 5)
		}
		//		if count > 10 {
		//			log.Printf("id%v:breaking connection [for test], bye", c.runid)

		//			c.NetConn.Close()
		//			break
		//		}
	}
}

func (c *Link) readLoop() {
	defer func() {
		if r := recover(); r != nil {
			close(c.exitChan)
			log.Errorf("recover from err, %v", r)
		}
	}()

	for {
		content, err := c.Reader.ReadString('\n')
		if err != nil {
			log.Printf("id %v: read string failed, %v", c.runid, err)
			close(c.exitChan)
			c.NetConn.Close()
			break
		} else {
			log.Printf("id %v <---: %v", c.runid, content)
		}
	}
}
