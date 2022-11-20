package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server listens at 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		ln := scan.Text()
		method := strings.Fields(ln)

		if method[0] == "GET" {
			message := "<h1>Hi from server</h1>"
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(message))
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			fmt.Fprintln(conn, message)
			fmt.Println(method[0])
			break
		}
	}

	defer conn.Close()
}
