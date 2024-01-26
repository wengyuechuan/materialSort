package materials

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	materials := r.Group("/materials")
	{
		materials.GET("", HelloMaterials)
		materials.GET("/category", GetMaterialsCategory)
		materials.POST("/category", AddMaterialsCategory)
		materials.PUT("/category", UpdateMaterialsCategory)
		materials.DELETE("/category", DeleteMaterialsCategory)
		materials.GET("/category/all", GetMaterialsCategoryAll)
		materials.GET("/category/:category_id", GetMaterialsCategoryById)
	}
}
