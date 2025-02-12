package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func figas() {
	ln, err := net.Listen("tcp", ":8081")

	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			c.Write([]byte("Hello, world!"))
			fmt.Println("Conexão estabelecida com: ", c.RemoteAddr())
		}(conn)

	}
}

func figas2(c net.Conn) {
	defer c.Close()
	fmt.Println("Conexão com: ", c.RemoteAddr())

	tcpConn, _ := c.(*net.TCPConn)
	tcpConn.SetKeepAlive(true)
	tcpConn.SetKeepAlivePeriod(2 * time.Minute)

	//leitor de comandos

	buffer := make([]byte, 1024)

	n, _ := c.Read(buffer)

	message := string(buffer[:n])

	switch message {
	case "PING":
		c.Write([]byte("PONG\n"))
	case "TIME":
		c.Write([]byte(time.Now().Format(time.RFC3339) + "\n"))
	default:
		c.Write([]byte("Comando desconhecido \n "))
	}
}

func main() {
	/*
		ln, err := net.Listen("tcp", ":8080")

		if err != nil {
			fmt.Println("Erro ao criar o listener:", err)
			return
		}

		defer func() { _ = ln.Close() }()

		fmt.Printf("Listener criado com sucesso %q\n", ln.Addr())

	*/
	//figas()

	ln, err := net.Listen("tcp", "[::]:8081")
	if err != nil {
		log.Fatal("Erro ao iniciar o sevidor", err)
	}

	defer ln.Close()

	fmt.Println("Servidor criado com sucesso ")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("Encerrando o servido: ")
		ln.Close()
		os.Exit(0)
	}()

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal("Erro na conexão ", err)
			continue
		}
		go figas2(conn)
	}
}
