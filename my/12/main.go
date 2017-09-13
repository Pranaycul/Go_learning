package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	In, err := net.Listen("tcp", ":9090")

	if err != nil {
		panic(err)
	}

	defer In.Close()

	for {
		conn, err := In.Accept()
		if err != nil {
			panic(err)
		}
		io.WriteString(conn, fmt.Sprint("Hello world \n", time.Now(), "\n"))
		conn.Close()
	}
}
