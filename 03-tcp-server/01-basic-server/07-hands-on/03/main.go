package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// create a multiplexer (mux, servemux, router, server, http mux, ...) so that your server responds to different URIs and METHODS.
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	var path string

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		path = strings.Fields(line)[1]
		fmt.Println(path)

		break
	}

	msg := "Wrong path"
	if path == "/hi" {
		msg = "Hey~"
	}

	respond(conn, msg)

}

func respond(conn net.Conn, msg string) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>` + msg + `</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
