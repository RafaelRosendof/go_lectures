package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func send_message(message string, conn net.Conn) net.Conn {

	conn.Write([]byte(message))
	reply := make([]byte, 1024)

	n, _ := conn.Read(reply)

	fmt.Println("Resposta do servidor: ", string(reply[:n]))

	return conn
}

func Connect_client() {

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error while connecting to server: ", err)
		return
	}

	defer conn.Close()

	go func() {
		for {
			reply := make([]byte, 1024)
			n, _ := conn.Read(reply)

			if err != nil {
				fmt.Println("\nConexão perdida com o servidor.")
				return
			}

			fmt.Printf("\r%s\nDigite sua mensagem: ", string(reply[:n]))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Digite sua mensagem: ")
		if !scanner.Scan() {
			break
		}

		text := scanner.Text()
		if text == "quit" {
			return
		}

		conn.Write([]byte(text))
	}
}
