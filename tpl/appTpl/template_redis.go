package appTpl

const TplRedis  =`package testFunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moka-mrp/sword-core/rds"
	. "{{.ModuleName}}/app/common/controllers"
	"time"
)

//http://localhost:9999/v1/test/redis
func Redis(c *gin.Context) {
	pools:=rds.GetRedis()
	//默认池子的使用

	_,err:=pools.Set("sword",time.Now().Unix())
	if err!=nil{
		fmt.Println(err)
	}
	name,err:=pools.Get("sword")
	if err !=nil{
		fmt.Println(err)
	}
	//另外池子的使用
	_,err=pools.AliasSet("center","sword",time.Now().Unix())
	if err !=nil{
		fmt.Println(err)
	}
	name2,err:=pools.AliasGet("center","sword")
	if err !=nil{
		fmt.Println(err)
	}
	Success(c,name+name2)
}`