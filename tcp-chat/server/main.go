package main

import (
	"fmt"
	"log"
	"net"
)

type connInfo struct {
	id int
}

var connections map[net.Conn]*connInfo

func init() {
	connections = make(map[net.Conn]*connInfo)
}

func handle(conn net.Conn, info *connInfo) {
	defer conn.Close()

	// Listen connection Data
	bs := make([]byte, 1024)
	for {
		read, err := conn.Read(bs)
		if err != nil {
			log.Printf("[%d] %s\n", info.id, err.Error())
			delete(connections, conn)
			return
		}

		msg := string(bs[:read])
		log.Printf("[%d] %s\n", info.id, msg)

		if msg == "ping" {
			_, err = conn.Write([]byte("[pong]"))
			if err != nil {
				log.Printf("[%d] %s\n", info.id, err.Error())
			}
			return
		}

		go broadcast(fmt.Sprintf("[%d] %s", info.id, msg), conn)
	}
}

// send message to all
func broadcast(msg string, from net.Conn) {
	for conn, info := range connections {
		if conn == from {
			continue
		}

		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Printf("[%d] %s\n", info.id, err.Error())
			delete(connections, conn)
			continue
		}
	}
}

func main() {
	// create tcp address
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8787")
	if err != nil {
		log.Fatalln(err)
	}

	// listen for tcp connections
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	log.Println("Server is running...")

	// accept tcp connections
	connCount := 0
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("[Failed]", conn)
			continue
		}
		connCount++

		info := &connInfo{connCount}
		connections[conn] = info
		log.Println("[New]", *info, conn.RemoteAddr().String())

		go handle(conn, info)
	}
}
