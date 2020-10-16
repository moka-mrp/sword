package appTpl

const TplServerRecovery  = `package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moka-mrp/sword-core/kernel/server"
	"github.com/moka-mrp/sword-core/log/logger"
	"{{.ModuleName}}/app/common/constants"
	"runtime/debug"
)

//统一的异常捕获中间件
//参照gin.Recovery()自带的中间件进行改写的额
//@author sam@2020-08-08 14:47:57
func ServerRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//记录到日志
				logger.Error(err)
				//本地开发 debug 模式开启时输出错误信息到shell
				if server.GetDebug(){
					fmt.Println(string(debug.Stack()))
				}
				//统一的响应:如果服务端发生错误或者异常，接口返回数据如下:
				c.JSON(200, gin.H{
					"errcode":constants.CodeSystemError ,
					"errmsg":"system error",
				})
				c.Abort()
			}
		}()
		//before request
		c.Next()
		//after request
	}
}
`