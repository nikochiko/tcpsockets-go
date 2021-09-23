# tcpsockets-go

A play project where I made a TCP client, server, and loadtesting tool for learning about TCP socket programming in Go.

I used https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/tcp_sockets.html as reference, and improvised on some parts (loadtest/loadtest.go for example, and supporting HTTP/1.1 HEAD requests by reading one byte at a time). 

## How to use:

* First run `go build`. This will create a `tcpsockets-go` binary in your local filesystem
* To make a HTTP HEAD request with raw TCP:
```shell
./tcpsockets-go client www.google.com:80
```
* To run the server on port 1200:
```shell
# multi-threaded by default
# ./tcpsockets-go server :$port
./tcpsockets-go server :1200

# for a single threaded server:
./tcpsockets-go server :1200 s
```
* To loadtest a local server:
```shell
# for 100 concurrent requests
./tcpsockets-go loadtest :1200 100
```
