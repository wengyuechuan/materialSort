package main

import (
	"github.com/gin-gonic/gin"
	"materialsSort/service/admin"
	"materialsSort/service/materials"
	"materialsSort/service/user"
	"materialsSort/service/welcome"
)

func InitRouter(r *gin.Engine) {
	welcome.Init(r)
	user.Init(r)
	materials.Init(r)
	admin.Init(r)
}
