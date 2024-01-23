package welcome

import "github.com/gin-gonic/gin"

func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Welcome to material-sort",
	})
}
