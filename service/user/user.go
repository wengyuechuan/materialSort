package user

import "github.com/gin-gonic/gin"

func HelloUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Hello User",
	})
}
