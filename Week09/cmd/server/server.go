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
	listener     *net.TCPListener
	responseChan chan *encoding.Response
}

var serverWriteLock sync.Mutex

func newServer() (server *Server, err error) {
	tcp, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8989")
	if err != nil {
		return nil, errors.Wrap(err, "resolve rtcp address")
	}
	listener, err := net.ListenTCP("tcp", tcp)
	if err != nil {
		return nil, errors.Wrap(err, "error tcp listen")
	}
	responseChan := make(chan *encoding.Response, 10)
	return &Server{listener: listener, responseChan: responseChan}, nil
}

func (s *Server) start() error {
	for {
		conn, err := s.listener.AcceptTCP()
		if err != nil {
			return errors.Wrap(err, "listen accept tcp")
		}
		go s.listen(conn)
		go s.send(conn)
	}
}

func (s *Server) listen(conn *net.TCPConn) {
	for {
		r, err := encoding.Reader(conn)
		if err != nil {
			fmt.Println(errors.Wrap(err, "listen err tcp"))
		}
		if r.Payload == "close" {
			s.responseChan <- r
			conn.Close()
			close(s.responseChan)
			break
		}
		s.responseChan <- r
	}
}

func (s *Server) send(conn *net.TCPConn) {
	for d := range s.responseChan {
		if d == nil {
			break
		}
		encoding.Write(&encoding.Response{Serial: d.Serial, Payload: d.Payload}, conn, &serverWriteLock)
	}
}

func (s *Server) closeListener() {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP)
	<-ch
	s.listener.Close()
}

func main() {
	ser, err := newServer()
	if err != nil {
		panic(err)
	}
	err = ser.start()
	if err != nil {
		fmt.Print("error", err)
	}
	go ser.closeListener()
}
