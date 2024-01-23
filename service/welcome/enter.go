package welcome

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	group := r.Group("/welcome")
	{
		group.GET("", Welcome)
	}
}
