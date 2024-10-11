package p2p_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/rcarvalho-pb/distributed-filesystem/p2p"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"

	tr := NewTCPTransport(listenAddr)

	assert.Equal(t, listenAddr, tr.GetListenAddress())

	assert.Nil(t, tr.ListenAndAccept())
}
