package main

import (
	"fmt"
	"log"
	"net"
	"strings"
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
		} else if v.typ != ARRAY_TEXT {
			log.Fatal("Invalid Type. Expected Array")
		}
		command := strings.ToLower(v.array[0].bulk)
		args := v.array[1:]

		handler, ok := handlers[command]
		response := Value{typ:STRING_TEXT, str:""}
		if ok {
			response = handler(args)
		}

		w := NewWriter(conn)
		_, e := w.write(response)
		if e != nil {
			log.Fatal(e)
		}
	}
}


