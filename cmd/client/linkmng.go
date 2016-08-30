package main

import (
	"sync"
	"time"

	"github.com/CodisLabs/codis/pkg/utils/log"
)

type LinkMng struct {
	linkid     int
	lock       sync.Mutex
	remoteAddr string
	links      map[int]*Link

	stats map[int]int
}

func newLinkMng(remoteAddr string) *LinkMng {
	return &LinkMng{
		remoteAddr: remoteAddr,
		links:      make(map[int]*Link),
		stats:      make(map[int]int),
	}
}

func (l *LinkMng) addStats(id int, count int) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.stats[id] = count
}

func (l *LinkMng) getStats() {
	var oldsum int = 0
	for {

		var sum int = 0
		l.lock.Lock()
		for _, count := range l.stats {
			sum += count
		}
		l.lock.Unlock()

		diff := sum - oldsum

		log.Infof("diff = %20d, sum = %20d, oldsum=%20d", diff, sum, oldsum)
		oldsum = sum

		<-time.After(time.Second)
	}
}

func (l *LinkMng) startClient() {
	l.lock.Lock()
	var id = l.linkid
	l.linkid += 1

	log.Infof("got id %d", id)

	link, err := NewLinkDial(id, l.remoteAddr, l)
	l.links[id] = link

	l.lock.Unlock()

	if err != nil {
		log.InfoErrorf(err, "build connection failed")
	} else {
		link.Run()
	}
}

func (l *LinkMng) Main() {
	for i := 0; i < int(*clientCount); i++ {
		log.Infof("start client %d", i)
		go l.startClient()
	}

	l.getStats()
}
