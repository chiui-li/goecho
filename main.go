package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	const resHeader = "HTTP/1.1 200 OK\r\nContent-Type: text/plain; charset=utf-8\r\nConnection: close\r\n\r\n"
	reqTime := "time:" + time.Now().Format("2006-01-02 15:04:05") + "\n"
	addr := "TCP: " + conn.RemoteAddr().String() + "\n"
	content := bufio.NewReader(conn)
	conn.Write([]byte(resHeader))
	conn.Write([]byte(reqTime))
	conn.Write([]byte(addr + "\n"))
	response := []byte{}
	reqHeaderEndIndex := 0
	bodyIndex := 0
	for {
		responseLine, err := content.ReadBytes('\n')
		if err != nil {
			break
		}
		response = append(response, responseLine...)
		if len(response) >= 4 && bodyIndex == 0 && bytes.Equal(response[len(response)-4:], []byte("\r\n\r\n")) {
			bodyIndex = len(response)
			reqHeaderEndIndex = bodyIndex - 4
			conn.Write(response[:reqHeaderEndIndex])
			break
		}
	}
	tcp, _ := conn.(*net.TCPConn)
	tcp.CloseWrite()
	log.Println(addr)
	log.Println(string(response))
}

func main() {
	server, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, _ := server.Accept()
		go handleConn(conn)
	}
}
