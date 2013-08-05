pooler
===============

Socket acceptor pool for `golang`.

Usage
==============

```go
type Client struct {
	Conn net.Conn
	Quit chan bool
}

func test_handler(conn *Client, data []byte) {
    // do something with data
    ....
    // close connection
	conn.close()
}

...
start_tcp_pool("localhost:8080", 1000, test_handler)
...
```

TODO
==============

  * add ssl

Contribute
==============

  * Fork https://github.com/0xAX/pooler
  * Make changes
  * Send pull request

Author
========

[@0xAX](https://twitter.com/0xAX).
