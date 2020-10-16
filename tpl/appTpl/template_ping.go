package appTpl

const TplPing  =`package healthFunc

import (
	"github.com/gin-gonic/gin"
	"github.com/moka-mrp/sword-core/rds"
	"{{.ModuleName}}/app/common/controllers"
	"net/http"
)

// 健康检查
//@todo 目前只检测redis
//@author sam@2020-08-07 14:25:43
func Ping(ctx *gin.Context) {
	_, err := rds.GetRedis().RawCommand("PING")
	if err != nil {
		controllers.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	controllers.Success(ctx, "")
	return
}`