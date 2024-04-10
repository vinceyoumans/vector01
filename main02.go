package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"net"
	"time"
)

func loggingServer() {
	listener, err := net.Listen("tcp", "localhost:50002")
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("failed to accept connection: %v\n", err)
			}
			go handlerConn(conn)
		}
	}()
}

func handlerConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("faild to read from conn: %v\n", err)
			break
		}
		fmt.Println("In loggingServer: ", msg)
	}
}

func main() {
	go loggingServer()

	conn, err := net.Dial("tcp", "localhost:50002")
	if err != nil {
		log.Fatal(err.Error())
	}

	logger := slog.New(slog.NewJSONHandler(conn, nil))
	logger.Info("Test the log to the connection")
	logger.Info("It works well")

	time.Sleep(5 * time.Second)
}
