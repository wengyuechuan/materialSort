package materials

import "github.com/gin-gonic/gin"

func HelloMaterials(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Hello Materials",
	})
}
