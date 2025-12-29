package main

import (
	"flag"
	"fmt"
	"tcp/tcp"
)

func main() {

	fmt.Println("Starting tcp server: ")

	mode := flag.String("mode", "server", "Define se é 'server' ou 'client'")
	flag.Parse()

	if *mode == "server" {
		tcp.Setup()
	} else {
		fmt.Println("Starting client")
		tcp.Connect_client()
		//RunClient()
	}
}
