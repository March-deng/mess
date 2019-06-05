package web

import (
	"net"
	"sync"
)

type epoll struct {
	fd          int
	connections map[int]net.Conn
	lock        *sync.RWMutex
}

func NewEpoll() (*epoll, error) {
	// fd, err := unix.
	return nil, nil
}
