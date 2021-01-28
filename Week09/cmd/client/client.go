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

//Client ...
type Client struct {
	conn *net.TCPConn
	seq  uint32
	data chan string
}

//NewClient ..
func NewClient() (cli *Client, err error) {
	tcp, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8989")
	if err != nil {
		return nil, errors.Wrap(err, "client resolve tcp")
	}
	conn, err := net.DialTCP("tcp", nil, tcp)
	if err != nil {
		return nil, err
	}
	dataChan := make(chan string, 10)
	return &Client{conn: conn, seq: 0, data: dataChan}, nil
}

func (c *Client) start() {
	go c.response()
	go c.request()

}

func (c *Client) response() {
	for {
		r, err := encoding.Reader(c.conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if r.Payload == "close" {
			c.conn.Close()
			break
		}
		fmt.Println("client :", r. )
	}
}

func (c *Client) request() {
	for d := range c.data {
		err := encoding.Write(&encoding.Response{Serial: c.seq, Payload: d}, c.conn, &clientWriteLock)
		if err != nil {
			fmt.Println(err)
		}
		seq++
	}
}

func main() {
	cli, err := NewClient()
	if err != nil {
		panic(err)
	}
	cli.start()
	for {
		paylod := "hi"
		cli.data <- paylod
		if paylod == "close" {
			fmt.Println("recieve close command")
			break
		}
	}
}
