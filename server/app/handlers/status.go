package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/imriyaya/mango/app/type"
)

func StatusHandler(server *_type.Server) {
	server.Router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
