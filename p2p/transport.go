package p2p

import "net"

// Peer is an interface that represents the remote node
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// Transport is anything that handles the communication
// between the nodes in the network. THis can be of the
// form (TCP, UDP, websockets...)
type Transport interface {
	Addr() string
	Dail(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
