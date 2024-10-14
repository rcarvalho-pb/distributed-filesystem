package p2p_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/rcarvalho-pb/distributed-filesystem/p2p"
)

func TestTCPTransport(t *testing.T) {
	tcpOpts := TCPTransportOpts{
		ListenAddress: ":3000",
		Handshaker:    NOOPHandshakeFunc,
		// Decoder: ,
	}

	tr := NewTCPTransport(tcpOpts)

	assert.Equal(t, tcpOpts.ListenAddress, tr.GetListenAddress())

	assert.Nil(t, tr.ListenAndAccept())
}
