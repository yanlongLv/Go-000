package encoding

import (
	"encoding/binary"
	"io"
	"net"
	"sync"

	"github.com/pkg/errors"
)

//Response ...
type Response struct {
	Serial  uint32
	Payload string
}

//Reader ..
func Reader(conn *net.TCPConn) (*Response, error) {
	ret := &Response{}
	buffer := make([]byte, 4)
	if _, err := io.ReadFull(conn, buffer); err != nil {
		return nil, errors.Wrap(err, "Reader response")
	}
	length := binary.BigEndian.Uint32(buffer)
	if _, err := io.ReadFull(conn, buffer); err != nil {
		return nil, errors.Wrap(err, "big end read full")
	}
	ret.Serial = binary.BigEndian.Uint32(buffer)
	payloadBytes := make([]byte, length-4)
	if _, err := io.ReadFull(conn, payloadBytes); err != nil {
		return nil, errors.Wrap(err, "read payload error")
	}
	ret.Payload = string(payloadBytes)
	return ret, nil
}

//Write ..
func Write(r *Response, conn *net.TCPConn, lock *sync.Mutex) error {
	payloadbuf := []byte(r.Payload)
	serialbuf := make([]byte, 4)
	binary.BigEndian.PutUint32(serialbuf, r.Serial)
	length := uint32(len(payloadbuf) + len(serialbuf))
	lengthByte := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthByte, length)
	_, err := conn.Write(lengthByte)
	if err != nil {
		return errors.Wrap(err, "lengthByte conn write err")
	}
	_, err = conn.Write(payloadbuf)
	if err != nil {
		return errors.Wrap(err, "payloadbuf conn write err")
	}
	_, err = conn.Write(serialbuf)
	if err != nil {
		return errors.Wrap(err, "serialbuf conn write err")
	}
	return nil
}
