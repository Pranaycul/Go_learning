package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":9090")

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}
