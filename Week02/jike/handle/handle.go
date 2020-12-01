package handle

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Go-000/Week02/jike/db"
	"github.com/gin-gonic/gin"
)

// Handle Client server
type Handle struct {
	Client db.ServiceInterface
}

// NewHandleClient handle client
func NewHandleClient(client db.ServiceInterface) *Handle {
	return &Handle{Client: client}
}

// GetAllUser get all users
func (h *Handle) GetAllUser(c *gin.Context) {
	users, err := h.Client.GetUsers()
	if err != nil {
		log.Printf("get all users error %v\n", err)
	}
	fmt.Printf("%v", users)
	c.JSON(200, gin.H{
		"data": users,
	})

}

// GetUserNameByID get user inforation
func (h *Handle) GetUserNameByID(c *gin.Context) {
	ID := c.Param("Id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		log.Printf("get users info by userId strconv atoi id %v\n", err)
	}
	name, err := h.Client.GetUserNameByID(id)
	if err != nil {
		log.Printf("get users info by userId error %v\n", err)
	}
	c.JSON(200, gin.H{
		"data": name,
	})
}
