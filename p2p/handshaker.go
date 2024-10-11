package p2p

type HandshakeFunc func(any) error

func NOOPHandshakeFunc(any) error {
	return nil
}
