package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Go-000/Week09/internal/encoding"
	"github.com/pkg/errors"
)

//Server ..
type Server struct {
	conn         *net.TCPConn
	responseChan chan string
}

var serverWriteLock sync.Mutex

func server() error {
	tcp, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8980")
	if err != nil {
		return errors.Wrap(err, "resolve rtcp address")
	}
	listener, err := net.ListenTCP("tcp", tcp)
	if err != nil {
		return errors.Wrap(err, "error tcp listen")
	}
	defer listener.Close()
	go closeListener(listener)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			return errors.Wrap(err, "listen accept tcp")
		}
		dataChan := make(chan *encoding.Response)
		go listen(conn, dataChan)
		go send(conn, dataChan)
	}
}

func listen(conn *net.TCPConn, dataChan chan *encoding.Response) {
	for {
		r, err := encoding.Reader(conn)
		if err != nil {
			fmt.Println(errors.Wrap(err, "listen err tcp"))
		}
		if r.Payload == "close" {
			dataChan <- r
			conn.Close()
			close(dataChan)
			break
		}
		dataChan <- r
	}
}

func send(conn *net.TCPConn, dataChan chan *encoding.Response) {
	for d := range dataChan {
		if d == nil {
			break
		}
		encoding.Write(&encoding.Response{Serial: d.Serial, Payload: d.Payload}, conn, serverWriteLock)
	}
}

func closeListener(listener *net.TCPListener) {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP)
	<-ch
	listener.Close()
}
