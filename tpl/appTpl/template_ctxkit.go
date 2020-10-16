package appTpl

const TplCtxkit  = `package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moka-mrp/sword-core/http/ctxkit"
	"github.com/moka-mrp/sword-core/kernel/server"
	"{{.ModuleName}}/config"
)

//共享中间件以及控制器的一些公共变量，尤其是一些配置信息可以在此处共享出来额
//@todo 好处是可以追踪到同一请求产生的所有流水记录
func GenContextKit(c *gin.Context) {

	ctxkit.SetClientId(c,c.ClientIP())
	ctxkit.SetServerId(c,c.Request.RemoteAddr)
	ctxkit.SetHost(c,c.Request.Host)
	ctxkit.SetTraceId(c,c.GetHeader("X-Trace-Id"))
	ctxkit.SetJwtSecret(c,config.GetConf().Jwt.Secret)

	if server.GetDebug(){
		fmt.Printf("[GIN] %s|%s \r\n",c.ClientIP(),c.Request.Host)
	}

}`