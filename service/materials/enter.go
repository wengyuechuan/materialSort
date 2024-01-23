package materials

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	materials := r.Group("/materials")
	{
		materials.GET("", HelloMaterials)
	}
}
