package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Command struct {
	Fields []string
	Result chan string
}

func redisServer(commands chan Command) {
	var data = make(map[string]string)
	for cmd := range commands {
		if len(cmd.Fields) < 2 {
			cmd.Result <- "Expected at least 2 argument"
			continue
		}
		fmt.Println("GOT COMMAND", cmd)
		switch cmd.Fields[1] {
		case "GET":
			key := cmd.Fields[1]
			value := data[key]
			cmd.Result <- value
		case "SET":
			if len(cmd.Fields) != 3 {
				cmd.Result <- "EXPECTED VALUE"
				continue
			}
			key := cmd.Fields[1]
			value := cmd.Fields[2]
			data[key] = value
			cmd.Result <- ""
		case "DEL":
			key := cmd.Fields[1]
			delete(data, key)
			cmd.Result <- ""
		default:
			cmd.Result <- "INVALID COMMAND" + cmd.Fields[0] + "\n"
		}

	}

}

func handle(commands chan Command, conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	ln := scanner.Text()
	fs := strings.Fields(ln)
	result := make(chan string)
	commands <- Command{
		Fields: fs,
		Result: result,
	}
	io.WriteString(conn, <-result+"\n")
}

func main() {
	li, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln(err)

	}
	defer li.Close()

	commands := make(chan Command)
	go redisServer(commands)
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(commands, conn)
	}
}
