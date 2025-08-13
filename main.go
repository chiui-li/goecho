package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	resHeader := "HTTP/1.1 200 OK\r\nContent-Type: text/plain; charset=utf-8\r\nConnection: keep-alive\r\n\r\n"
	reqTime := "time:" + time.Now().Format("2006-01-02 15:04:05") + "\n"
	addr := "HTTP: " + conn.RemoteAddr().String() + "\n"
	content := bufio.NewReader(conn)
	log.Println(reqTime, addr)
	conn.Write([]byte(resHeader))
	conn.Write([]byte(reqTime))
	conn.Write([]byte(addr + "\n"))
	conn.Write([]byte(reqTime))
	content.WriteTo(conn)
	tcp, _ := conn.(*net.TCPConn)
	tcp.CloseWrite()
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
