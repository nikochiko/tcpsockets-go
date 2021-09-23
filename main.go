package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nikochiko/tcpsockets-go/client"
	"github.com/nikochiko/tcpsockets-go/loadtest"
	"github.com/nikochiko/tcpsockets-go/server"
	"github.com/nikochiko/tcpsockets-go/util"
)

const (
	defaultServerMultithreadFlag = true
	defaultLoadtestConcurrency = 100
)

func main() {
	if len(os.Args) < 3 {
		err := fmt.Errorf("usage: %s [client|server|loadtest] [<host>:<port>|<port>|<host:port>] [m|s|<concurrency>]", os.Args[0])
		util.CheckError(err)
	}

	switch app := os.Args[1]; app {
	case "client":
		service := os.Args[2]
		fmt.Printf("Starting client app with arg: %s\n", service)
		client.MakeHEADRequest(service)
	case "server":
		service := os.Args[2]
		fmt.Printf("Starting server app with arg: %s\n", service)
		multithreadFlag := defaultServerMultithreadFlag
		if len(os.Args) > 3 && os.Args[3] == "s" {
			multithreadFlag = false
		}
		server.StartListening(service, multithreadFlag)
	case "loadtest":
		service := os.Args[2]
		concurrency := defaultLoadtestConcurrency
		if len(os.Args) > 3 {
			var err error
			concurrency, err = strconv.Atoi(os.Args[3])
			util.CheckError(err)
		}
		loadtest.LoadTest(service, concurrency, []byte{})
	default:
		err := fmt.Errorf("unrecognized app: %s", app)
		util.CheckError(err)
	}
}
