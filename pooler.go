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
type Handler func(conn *Client, data []byte)

//
// Client structure
//
type Client struct {
	Conn net.Conn
	Quit chan bool
}

//
// Close connection
//
func (c *Client) close() {
	c.Conn.Close()
	c.Quit <- false
}

//
// Listener state. Holds acceptors count
//
func listenerState(acceptorsCount *int, client *Client) {
	// catch 'connection close'
	for {
		select {
		case <-client.Quit:
			*acceptorsCount--
			return
		}
	}
}

//
// Handle tcp connection. receive data and send to handler
//
func handleTcpConnection(conn net.Conn, handler Handler, acceptorsCount *int, client *Client) {
	//
	// increase acceptor number
	//
	*acceptorsCount = *acceptorsCount + 1

	//
	// start handle new connection
	//
	for {
		// read data
		line, err := bufio.NewReader(conn).ReadBytes('\n')

		//
		// check error in reading data
		//
		if err != nil {
			if err == io.EOF {
				// decrease clients number
				*acceptorsCount = *acceptorsCount - 1
				// close connection
				client.close()
			}

			return
		}

		//
		// send data to handler
		//
		handler(client, line)
	}
}

//
// Start new tcp pool
//
func tcp_pool(addr string, acceptorsNum int, handler Handler) {

	acceptors := new(int)

	*acceptors = 0

	// start listen
	listener, err := net.Listen("tcp", addr)

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
			fmt.Println("acceptors in listener", *acceptors)
			//
			// accept new connection
			//
			connection, err := listener.Accept()

			//
			// check connection error
			//
			if err != nil {
				println("Error accept:", err.Error())
			}

			defer connection.Close()

			// create new client
			client := &Client{connection, make(chan bool)}

			// start new listener state
			go listenerState(acceptors, client)
			// handle new tcp connection
			go handleTcpConnection(connection, handler, acceptors, client)
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
