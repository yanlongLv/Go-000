package main

import (
	"github.com/alwayGo/newProject/api/v1/news"
	"github.com/alwayGo/newProject/internal/biz"
	service "github.com/alwayGo/newProject/internal/service/news"
	"google.golang.org/grpc"
)

func initNewServer (s *grpc.Server){
	initCli:= biz.NewClient()
	server:=service.NewserviceClient(initCli)
	news.RegisterNewsServer(s,server)
}
