package main

import (
	"log"

	"github.com/nadal07/DistributedFileStorage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
	// fmt.Println("We Gucci")
}
