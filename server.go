package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"strings"
)

func handle(conn net.Conn) {
	log.Printf("received request from %s\n", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	defer conn.Close()
	_, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}

	split := strings.Split(string(buf), " ")
	if len(split) < 2 {
		res := response{status: http.StatusBadRequest}
		conn.Write(res.serializeToBytes())
		return
	}

	path := split[1]
	res := resolveRoute(path)

	conn.Write(res.serializeToBytes())
}

func startServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	defer listener.Close()

	log.Printf("web server listening on port %s\n", port)

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Fatal(err)
			} else {
				log.Print(err)
				continue
			}
		}

		go handle(conn)
	}
}
