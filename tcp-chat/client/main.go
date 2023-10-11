package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// create tcp address
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8787")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// dial tcp address
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Listen connection Data
	go func() {
		bs := make([]byte, 1024)
		for {
			read, err := conn.Read(bs)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			msg := string(bs[:read])
			fmt.Printf("%s\n", msg)
		}
	}()

	// input text
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			conn.Write([]byte(scanner.Text()))
		}
		if err = scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}
