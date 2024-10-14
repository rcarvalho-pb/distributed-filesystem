package p2p

type Peer interface {
	Close() error
}

// lida com a conexão entre dois pontos
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
