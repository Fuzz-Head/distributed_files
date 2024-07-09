package p2p

// HandshakerFUnc ..?
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
