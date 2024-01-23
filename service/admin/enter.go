package admin

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.GET("", HelloAdmin)
	}
}
