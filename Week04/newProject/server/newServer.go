package server

import (
	"google.golang.org/grpc"
)

type ServerClient struct {
	*grpc.Server
}

func NewServerClient (server *grpc.Server) *ServerClient {
	srv := grpc.NewServer()
	return &ServerClient{srv}
}


