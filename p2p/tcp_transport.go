package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.Mutex
	peers map[net.Addr]Peer
}

func (t *TCPTransport) GetListenAddress() string {
	return t.listenAddress
}

func NewTCPTransport(listenAddr string) TCPTransport {
	return TCPTransport{
		listenAddress: listenAddr,
	}
}
