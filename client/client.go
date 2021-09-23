package client

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/nikochiko/tcpsockets-go/util"
)

type CustomReader struct {
	Stream io.Reader
}

// ReadUntilNewline reads len(b) bytes from reader or until a sequence of bytes seq is encountered
func (cr *CustomReader) ReadUntil(seq, buf []byte) (int, error) {
	characterQueue := make([]byte, len(seq))
	singleCharBuf := make([]byte, 1)

	for i := 0; i < len(buf); i++ {
		_, err := cr.Stream.Read(singleCharBuf)
		if err != nil {
			return 0, err
		}

		b := singleCharBuf[0]
		buf[i] = b

		characterQueue = addByteToQueue(characterQueue, b)

		if bytes.Equal(characterQueue, seq) {
			return i + 1, nil
		}
	}

	return len(buf), nil
}

func addByteToQueue(q []byte, b byte) []byte {
	for i := len(q) - 1; i >= 0; i-- {
		if i == 0 {
			q[i] = b
		} else {
			q[i] = q[i-1]
		}
	}

	return q
}

func MakeHEADRequest(service string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	util.CheckError(err)

	parts := strings.SplitN(service, ":", 2)
	host := parts[0]

	fmt.Printf("TCPAddr: %v\n", tcpAddr)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	util.CheckError(err)

	body := fmt.Sprintf("HEAD / HTTP/1.1\r\nHost: %s\r\nUser-Agent: goclient/1.1\r\nAccept: application/json\r\n\r\n", host)
	_, err = conn.Write([]byte(body))
	util.CheckError(err)

	buf := make([]byte, 64)
	r := CustomReader{Stream: conn}

	for true {
		bytesRead, err := r.ReadUntil([]byte("\n"), buf)
		util.CheckError(err)

		if bytes.Equal(buf[:bytesRead], []byte("\r\n")) {
			break
		}

		fmt.Printf(string(buf[:bytesRead]))
	}

	fmt.Println("")

	conn.Close()

	os.Exit(0)
}
