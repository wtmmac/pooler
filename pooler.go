package pooler

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

//
// connection handler
//
type Handler func(conn net.Conn, data []byte)

//
// Handle tcp connection. receive data and send to handler
//
func handleTcpConnection(conn net.Conn, handler Handler, acceptorsCount *int) {
	for {
		// read data
		line, err := bufio.NewReader(conn).ReadBytes('\n')

		//
		// check error in reading data
		//
		if err != nil {
			if err == io.EOF {
				*acceptorsCount = *acceptorsCount - 1
			}

			return
		}

		handler(conn, line)
	}
}

//
// Start new tcp pool
//
func tcp_pool(addr string, acceptorsNum int, handler Handler) {
	acceptors := new(int)

	*acceptors = 1

	// start listen
	listener, err := net.Listen("tcp", ":8080")

	//
	// check error
	//
	if err != nil {
		fmt.Println("error listening:", err.Error())
		os.Exit(1)
	}

	for {
		//
		// check acceptors number
		//
		if *acceptors == acceptorsNum {
		} else {
			//
			// accept new connection
			//
			connection, err := listener.Accept()

			//
			// increase acceptor number
			//
			*acceptors = *acceptors + 1

			//
			// check connection error
			//
			if err != nil {
				println("Error accept:", err.Error())
			}

			defer connection.Close()

			go handleTcpConnection(connection, handler, acceptors)
		}
	}
}

//
// Start new tcp pool
//
func start_tcp_pool(addr string, acceptorsNum int, handler Handler) {
	// start new pool
	go tcp_pool(addr, acceptorsNum, handler)
}
