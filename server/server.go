package server

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/nikochiko/tcpsockets-go/util"
)

func StartListening(service string, multiThread bool) {
	fmt.Printf("Resolving TCP addr for %s\n", service)
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)

	util.CheckError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	util.CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		if multiThread {
			// multi-threaded
			go HandleClient(conn)
		} else {
			// single-threaded
			HandleClient(conn)
		}
	}
}

func HandleClient(conn net.Conn) {
	uniqueID := rand.Intn(1000)

	fmt.Printf("[%d] Got a request\n", uniqueID)

	daytime := time.Now().String()
	conn.Write([]byte(fmt.Sprintf("%s\r\n", daytime)))

	conn.Close()
	fmt.Printf("[%d] Completed the request\n", uniqueID)
}
