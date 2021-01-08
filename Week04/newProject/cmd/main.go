package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis,err:=net.Listen("tcp",":8080")
	if err!=nil {
		log.Fatal(err,"listen error")
	}
	s:=grpc.NewServer()
	initNewServer(s)
	if err:= s.Serve(lis); err!=nil {
		log.Fatal("failed to start service")
	}
}
