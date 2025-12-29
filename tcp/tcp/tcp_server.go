package tcp

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func Setup() {

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error while try to setup the tcp connection")
		return
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			continue
		}

		go HandleConnection(conn)
	}
}

var (
	clients = make(map[net.Conn]bool)
	mutex   = sync.Mutex{}
)

func HandleConnection(c net.Conn) {
	defer func() {
		mutex.Lock()
		delete(clients, c)
		mutex.Unlock()
		c.Close()
	}()

	mutex.Lock()
	clients[c] = true
	mutex.Unlock()

	if tcpConn, ok := c.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(2 * time.Minute)
	}

	c.Write([]byte("Welcome to figas server \n"))

	buffer := make([]byte, 1024)

	for {

		n, err := c.Read(buffer)
		if err != nil {
			return
		}

		message := string(buffer[:n])
		fmt.Printf("Received from %s: %s \n", c.RemoteAddr(), message)

		mutex.Lock()
		for cliente := range clients {
			if cliente != c {
				cliente.Write([]byte("\nAlguém disse: " + message + "\nDigite sua mensagem: "))
			}
		}
		mutex.Unlock()
	}
}
