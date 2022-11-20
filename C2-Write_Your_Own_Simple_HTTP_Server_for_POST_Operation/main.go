package main

import (
	"bufio"
	"encoding/json"
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
	i := 0
	var method string
	var body string
	for scan.Scan() {
		ln := scan.Text()

		if i == 0 {
			method = strings.Fields(ln)[0]
		}

		if i >= 10 {
			body += ln
		}

		if ln != "" && ln[len(ln)-1] != []byte(",")[0] && i > 10 {
			break
		}
		i++
	}

	if method == "POST" {
		body += "}"
		postHandler(conn, body)
	}
	defer conn.Close()
}

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func postHandler(conn net.Conn, body string) {
	req := Person{}

	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		sendResponse(conn, "500 Internal Server Error", "internal server error", "text/plain")
		return
	}

	res, err := json.Marshal(req)
	if err != nil {
		sendResponse(conn, "500 Internal Server Error", "internal server error", "text/plain")
		return
	}

	sendResponse(conn, "200 OK", string(res), "application/json")
}

func sendResponse(conn net.Conn, httpStatus, message, contentType string) {
	fmt.Fprintf(conn, "HTTP/1.1 %v\r\n", httpStatus)
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(message))
	fmt.Fprintf(conn, "Content-Type: %v\r\n", contentType)
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintln(conn, message)
}
