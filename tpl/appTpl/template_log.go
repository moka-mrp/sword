package appTpl

const TplLog  =`package testFunc

import (
	"github.com/gin-gonic/gin"
	"github.com/moka-mrp/sword-core/log/logger"
	. "{{.ModuleName}}/app/common/controllers"
)


//todo 提供支持json输出，方便elk采集
//http://localhost:9999/v1/test/log
//@author sam@2020-08-08 15:43:09
func Log(c *gin.Context) {
	//方式一
	logger.GetLogger().Info("sam is a good man.")
	//方式二
	logger.Info("sam is a good man2.")

	Success(c, "how to use logger")
	return
}`