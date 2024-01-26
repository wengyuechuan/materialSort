package materials

import (
	"github.com/gin-gonic/gin"
	"materialsSort/models"
	"materialsSort/models/res"
	"strconv"
)

// HelloMaterials 简单路由，用于判断 materials 路由是否正常
func HelloMaterials(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Hello Materials",
	})
}

// GetMaterialsCategory 获取分类
func GetMaterialsCategory(c *gin.Context) {
	// 利用c获取页码和每页数量
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")
	// 判断两个参数是否存在
	if pageNum == "" || pageSize == "" {
		pageNum = "1"
		pageSize = "10"
	}
	// 转换成int类型
	pageNumInt, err := strconv.Atoi(pageNum)
	if err != nil {
		pageNumInt = 1
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10
	}

	count, category, err := models.GetCategory(pageNumInt, pageSizeInt)
	if err != nil {
		c.JSON(200, gin.H{
			"code": res.SystemError,
			"msg":  res.ErrorMap[res.SystemError],
		})
		return
	}
	c.JSON(200, gin.H{
		"code": res.SuccessCode,
		"msg":  "获取分类成功",
		"data": gin.H{
			"count":    count,
			"category": category,
		},
	})
}

// AddMaterialsCategory 添加分类
func AddMaterialsCategory(c *gin.Context) {
	categoryName := c.PostForm("category_name")
	if categoryName == "" {
		c.JSON(200, gin.H{
			"code": res.ArgumentError,
			"msg":  res.ErrorMap[res.ArgumentError],
		})
		return
	}
	err := models.AddCategory(categoryName)
	if err != nil {
		c.JSON(200, gin.H{
			"code": res.SystemError,
			"msg":  res.ErrorMap[res.SystemError],
		})
		return
	}
	c.JSON(200, gin.H{
		"code": res.SuccessCode,
		"msg":  "添加分类成功",
	})
}

// UpdateMaterialsCategory 更新分类
func UpdateMaterialsCategory(c *gin.Context) {
	categoryId := c.PostForm("category_id")
	categoryName := c.PostForm("category_name")
	if categoryId == "" || categoryName == "" {
		c.JSON(200, gin.H{
			"code": res.ArgumentError,
			"msg":  res.ErrorMap[res.ArgumentError],
		})
		return
	}
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		c.JSON(200, gin.H{
			"code": res.ArgumentError,
			"msg":  res.ErrorMap[res.ArgumentError],
		})
		return
	}
	err = models.UpdateCategory(categoryIdInt, categoryName)
	if err != nil {
		c.JSON(200, gin.H{
			"code": res.SystemError,
			"msg":  res.ErrorMap[res.SystemError],
		})
		return
	}
	c.JSON(200, gin.H{
		"code": res.SuccessCode,
		"msg":  "更新分类成功",
	})
}

// DeleteMaterialsCategory 删除分类
func DeleteMaterialsCategory(c *gin.Context) {
	categoryId := c.PostForm("category_id")
	if categoryId == "" {
		c.JSON(200, gin.H{
			"code": res.ArgumentError,
			"msg":  res.ErrorMap[res.ArgumentError],
		})
		return
	}
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		c.JSON(200, gin.H{
			"code": res.ArgumentError,
			"msg":  res.ErrorMap[res.ArgumentError],
		})
		return
	}
	err = models.DeleteCategory(categoryIdInt)
	if err != nil {
		c.JSON(200, gin.H{
			"code": res.SystemError,
			"msg":  res.ErrorMap[res.SystemError],
		})
		return
	}
	c.JSON(200, gin.H{
		"code": res.SuccessCode,
		"msg":  "删除分类成功",
	})
}

// GetMaterialsCategoryAll 获取所有分类
func GetMaterialsCategoryAll(c *gin.Context) {
	count, category, err := models.GetCategoryAll()
	if err != nil {
		c.JSON(200, gin.H{
			"code": res.SystemError,
			"msg":  res.ErrorMap[res.SystemError],
		})
		return
	}
	c.JSON(200, gin.H{
		"code": res.SuccessCode,
		"msg":  "获取所有分类成功",
		"data": gin.H{
			"count":    count,
			"category": category,
		},
	})
}

// GetMaterialsCategoryById 根据分类id获取分类
func GetMaterialsCategoryById(c *gin.Context) {
	categoryId := c.Param("category_id")
	if categoryId == "" {
		c.JSON(200, gin.H{
			"code": res.ArgumentError,
			"msg":  res.ErrorMap[res.ArgumentError],
		})
		return
	}
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		c.JSON(200, gin.H{
			"code": res.ArgumentError,
			"msg":  res.ErrorMap[res.ArgumentError],
		})
		return
	}
	category, err := models.GetCategoryById(categoryIdInt)
	if err != nil {
		c.JSON(200, gin.H{
			"code": res.SystemError,
			"msg":  res.ErrorMap[res.SystemError],
		})
		return
	}
	c.JSON(200, gin.H{
		"code": res.SuccessCode,
		"msg":  "获取分类成功",
		"data": gin.H{
			"category": category,
		},
	})
}
