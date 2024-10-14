package main

import (
	"log"

	"github.com/rcarvalho-pb/distributed-filesystem/p2p"
)

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		Handshaker:    p2p.NOOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
