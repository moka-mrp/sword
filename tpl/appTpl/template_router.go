package appTpl

const TplRouter  = `package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moka-mrp/sword-core/http/middleware"
	"{{.ModuleName}}/app/api/controller/adminFunc"
	"{{.ModuleName}}/app/api/controller/healthFunc"
	"{{.ModuleName}}/app/api/controller/testFunc"
	"{{.ModuleName}}/app/common/middlewares"
    "{{.ModuleName}}/app/common/utils"
)

//分布式定时任务模块的路由
//todo 注意GenContextKit中间件比较特殊，它的工作在内部已经完成了，不需要执行回调()
//todo 按标准来写，既规范化同时也能避免路由冲突额
//@author  sam@2020-08-17 14:07:45
func RegisterRoute(router *gin.Engine) {
	//统一的中间件
	router.Use(middleware.Cors(),middlewares.ServerRecovery(),middlewares.GenContextKit,gin.Logger())
	//必备的路由
	router.GET("/", func(context *gin.Context) {
		context.String(200,utils.GetVersion())
	})
	router.GET("/ping", healthFunc.Ping) // 健康检查
    //v1接口路由
	api := router.Group("/v1")
	//1.账户管理  sam@2020-08-17 14:28:29
	admin:= api.Group("/admin")
	{
		admin.POST("/login", adminFunc.Login)
		admin.Use(middlewares.JwtTokenVerify())
		admin.PUT("/account",adminFunc.AddAccount)
	}
	//2.测试管理  sam@2020-08-17 14:33:06
	test := api.Group("/test")
	{
		test.POST("/fast", testFunc.Fast)           //快速入门体验
		test.GET("/log", testFunc.Log)              //Log示例
		test.POST("/validator", testFunc.Validator) //Validator示例
		test.GET("/redis", testFunc.Redis)          //Redis示例
	}









}`