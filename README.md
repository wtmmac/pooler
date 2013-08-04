pooler
===============

Socket acceptor pool for `golang`.

Usage
==============

```go
func test_handler(conn net.Conn, data []byte) {
	fmt.Println(data)
}

func main() {
    // start new tcp pool
    start_tcp_pool(":8080", 1000, handler)
}
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