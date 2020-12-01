package routes

import (
	"github.com/Go-000/Week02/jike/db"
	"github.com/Go-000/Week02/jike/handle"
	"github.com/gin-gonic/gin"
)

// Route service route
func Route(r *gin.Engine, client db.ServiceInterface) {
	handle := handle.NewHandleClient(client)
	r.GET("/api/v1/users/list", handle.GetAllUser)
	r.GET("/api/v1/user/:Id", handle.GetUserNameByID)
}
