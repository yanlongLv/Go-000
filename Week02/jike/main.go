package main

import (
	"log"
	"time"

	"github.com/Go-000/Week02/jike/db"
	"github.com/Go-000/Week02/jike/error"
	"github.com/Go-000/Week02/jike/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.New()
	r.Use(error.Recover)
	serverClient := initClient()
	routes.Route(r, serverClient)
	r.Run(":8080")

}

func initClient() db.ServiceInterface {
	connect, err := db.NewConnect()
	if err != nil {
		log.Printf("connect error %v\n", err)
		time.Sleep(time.Second * 10)
		initClient()
	}
	return db.NewServerClient(connect)
}
