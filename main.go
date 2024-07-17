package main

import (
	"log"

	"github.com/Fuzz-Head/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcptransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		//TODO onPeer func
	}

	tcpTransport := p2p.NewTCPTransport(tcptransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	s := NewFileServer(fileServerOpts)
	tcpTransport.OnPeer = s.OnPeer

	return s
}
func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(s1.Start())
	}()

	s2.Start()
}

// func OnPeer(peer p2p.Peer) error {
// 	//fmt.Println("doing some logic with the peer outside of TCPTransport")
// 	peer.Close()
// 	return nil
// }

// func main() {
// 	tcpOpts := p2p.TCPTransportOpts{
// 		ListenAddr:    ":3000",
// 		HandshakeFunc: p2p.NOPHandshakeFunc,
// 		Decoder:       p2p.DefaultDecoder{},
// 		OnPeer:        OnPeer,
// 	}

// 	tr := p2p.NewTCPTransport(tcpOpts)

// 	go func() {
// 		for {
// 			msg := <-tr.Consume()
// 			fmt.Printf("%+v\n", msg)
// 		}
// 	}()

// 	if err := tr.ListenAndAccept(); err != nil {
// 		log.Fatal(err)
// 	}

// 	select {}
// }
