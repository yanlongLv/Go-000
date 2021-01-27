package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/Go-000/Week09/internal/encoding"
	"github.com/pkg/errors"
)

var clientWriteLock sync.Mutex
var seq uint32 = 0
var data = make(chan string, 10)

func client() error {
	tcp, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8989")
	if err != nil {
		return errors.Wrap(err, "client resolve tcp")
	}
	conn, err := net.DialTCP("tcp", nil, tcp)
	if err != nil {
		return err
	}
	go response(conn)
	go request(conn, data)
	return nil
}

func response(conn *net.TCPConn) {
	for {
		r, err := encoding.Reader(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if r.Payload == "close" {
			conn.Close()
			break
		}
		fmt.Println("client :", r.Payload)
	}
}

func request(conn *net.TCPConn, data chan string) {
	for d := range data {
		err := encoding.Write(&encoding.Response{Serial: seq, Payload: d}, conn, &clientWriteLock)
		if err != nil {
			fmt.Println(err)
		}
		seq++
	}
}

func main() {

	client()
}
