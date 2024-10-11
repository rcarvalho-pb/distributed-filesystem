package p2p

import (
	"fmt"
	"net"
	"sync"
)

type Temp struct{}

type TCPPeer struct {
	conn net.Conn

	outbound bool
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	handshaker    HandshakeFunc
	decoder       Decoder

	mu    sync.Mutex
	peers map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (t *TCPTransport) GetListenAddress() string {
	return t.listenAddress
}

func NewTCPTransport(listenAddr string) TCPTransport {
	return TCPTransport{
		handshaker:    NOOPHandshakeFunc,
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	if err := t.handshaker(conn); err != nil {
		conn.Close()
		return
	}

	msg := &Temp{}

	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("new incoming connection: %+v\n", conn)
			continue
		}
	}

}
