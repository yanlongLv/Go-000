package error

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recover middle ware
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover error: %v\n", r)
			c.JSON(http.StatusOK, gin.H{
				"code": 1001,
				"msg":  "系统维护中",
				"data": nil,
			})
			c.Abort()
		}
	}()
	c.Next()
}
