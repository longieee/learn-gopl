package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	portFlag := flag.Uint("port", 8000, "The port that the server listens on")
	flag.Parse()
	if *portFlag < 1024 || *portFlag > 65535 {
		log.Fatalf("port number %d out of range 1024..65535\n", *portFlag)
	}
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(int(*portFlag)))
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
