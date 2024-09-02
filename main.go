package main

import (
	"fmt"
	"net"
	"log"
)

func main() {
	fmt.Println("Listening on port 6379")
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		resp := newResp(conn)
		v, err := resp.read()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Data: %v", v)
		conn.Write([]byte("+OK\r\n"))
	}
}


