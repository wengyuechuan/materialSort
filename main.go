package main

import (
	"github.com/gin-gonic/gin"
	"materialsSort/core"
	"materialsSort/global"
)

func main() {
	core.InitConfig()              // 初始化配置文件
	global.Log = core.InitLogger() // 初始化日志
	global.DB = core.InitGorm()    // 初始化gorm

	r := gin.Default()
	InitRouter(r) //初始化路由
	addr := global.Config.System.Addr()
	err := r.Run(addr)
	if err != nil {
		global.Log.Fatal(err)
		return
	}
	global.Log.Infof("material-sort is running at %s", addr)
}
