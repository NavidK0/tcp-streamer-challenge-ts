package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", "127.0.0.1:64362")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("listening on addr: ", l.Addr())
	// Close the listener when the application closes.
	defer l.Close()
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	defer conn.Close()
	size := make([]byte, 2)
	_, err := conn.Read(size)

	n, err := strconv.Atoi(string(size))
	if n <= 12 || err != nil {
		errResp := "ERR: incorrect size"
		conn.Write([]byte(fmt.Sprintf("%d%s\n", len(errResp), errResp)))
		return
	}

	// Make a buffer to hold incoming data.
	buf := make([]byte, n)
	// Read the incoming connection into the buffer.
	_, err = conn.Read(buf)
	if err != nil {
		errResp := fmt.Sprintf("ERR: reading connection: %s", err.Error())
		conn.Write([]byte(fmt.Sprintf("%d%s\n", len(errResp), errResp)))
		return
	}
	if len(buf) <= 12 {
		errResp := "ERR: incorrect request length"
		conn.Write([]byte(fmt.Sprintf("%d%s\n", len(errResp), errResp)))
		return
	}

	cmd := string(buf[4:11])
	if cmd == "RANDNUM" {
		randLen := string(buf[12:n])
		n, err := strconv.Atoi(randLen)
		if err != nil {
			errResp := "ERR: not a number"
			conn.Write([]byte(fmt.Sprintf("%d%s\n", len(errResp), errResp)))
			return
		}
		if n > 10 {
			errResp := "ERR: number too large"
			conn.Write([]byte(fmt.Sprintf("%d%s\n", len(errResp), errResp)))
			return
		}

		numSlice := make([]string, 0, n)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < n; i++ {
			numSlice = append(numSlice, fmt.Sprintf("%d", rand.Intn(100-1)))
		}

		resp := strings.Join(numSlice, ",")
		respLen := fmt.Sprintf("%d", len(resp))
		if len(resp) < 9 {
			respLen = fmt.Sprintf("0%d", len(resp))
		}
		conn.Write([]byte(fmt.Sprintf("%s%s\n", respLen, resp)))
		return
	}

	errResp := "ERR: command not recognized"
	conn.Write([]byte(fmt.Sprintf("%d%s\n", len(errResp), errResp)))
}
