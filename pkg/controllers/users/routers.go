package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(router *gin.RouterGroup) {
	router.GET("/", Read)
}

func Read(c *gin.Context) {
	res := gin.H {
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Server is ready!",
		},
	}
	c.JSON(http.StatusOK, res)
}