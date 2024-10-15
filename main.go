package main

import (
	"fmt"
	"log"

	"github.com/rcarvalho-pb/distributed-filesystem/p2p"
)

func OnPeer(peer p2p.Peer) error {
	fmt.Printf("doing some logic here\n")
	peer.Close()
	return nil
}

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		Handshaker:    p2p.NOOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
