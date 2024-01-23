package user

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("", HelloUser)
	}
}
