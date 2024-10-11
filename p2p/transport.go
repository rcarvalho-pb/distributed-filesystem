package p2p

type Peer interface{}

// lida com a conexão entre dois pontos
type Transport interface {
	ListenAndAccept() error
}
