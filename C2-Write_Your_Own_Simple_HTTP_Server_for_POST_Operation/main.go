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
	i := 0
	var method string
	var body string
	for scan.Scan() {
		ln := scan.Text()

		if i == 0 {
			method = strings.Fields(ln)[0]
		}
		fmt.Println(ln, i)
		if ln == "" && i > 8 && scan.Scan() {
			// if scan.Scan() {
			// fmt.Println("here", scan.Scan())
			body = scan.Text()
			fmt.Println("body", body)

			// if method == "POST" {
			// 	fmt.Println("masuk method")
			// 	sendResponse(conn, "200 OK", body, "text/plain")
			// }
			break
			// }
		}

		i++
	}
	fmt.Println(method)
	fmt.Println(body)

	// if method == "POST" {
	// 	sendResponse(conn, "200 OK", body, "text/plain")
	// }

	defer conn.Close()
}

// func sendResponse(conn net.Conn, httpStatus, message, contentType string) {
// 	str := base64.StdEncoding.EncodeToString([]byte(message))

// 	fmt.Fprintf(conn, "HTTP/1.1 %v\r\n", httpStatus)
// 	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(str))
// 	fmt.Fprintf(conn, "Content-Type: %v\r\n", contentType)
// 	fmt.Fprint(conn, "\r\n")
// 	fmt.Fprintln(conn, str)
// }
