package admin

import "github.com/gin-gonic/gin"

func HelloAdmin(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Hello Admin",
	})
}
