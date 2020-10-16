package appTpl

const TplFast  = `package testFunc

import (
	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/app/common/constants"
	. "{{.ModuleName}}/app/common/controllers"
	"time"
)


//请求数据结构
type FastRequest struct {
	Name string `+"`json:\"name\" validate:\"required\" example:\"moka\"`"+`
	Url  string `+"`json:\"url\" validate:\"required,url\" example:\"https://www.baidu.com/\"`"+`
}
//返回数据结构
type FastResponse struct {
	Id   int64  `+"`json:\"id\" example:\"1\"`"+`
	Name string `+"`json:\"name\" example:\"moka\"`"+`
	Url  string `+"`json:\"url\" example:\"https://www.baidu.com/\"`"+`
}

//快速体验的一个案例
//curl -XPOST http://localhost:9999/v1/test/fast -H "Content-Type: application/json" -d  '{"name":"sam","url":"https://www.baidu.com"}'
//@author sam@2020-04-07 11:04:03
func Fast(c *gin.Context) {
	//(1)提取客户端传递的json字符串并检测参数
	request := new(FastRequest)
	err := GenRequest(c, request)
	if err != nil {
		Error(c,constants.CodeParamError) //参数检测失败之后统一返回错误即可
		return
	}
	//(2)组装返回数据
	response := new(FastResponse)
	response.Name = request.Name
	response.Url = request.Url
	response.Id = time.Now().Unix()
	//(3)成功返回即可
	Success(c, response)
	return
}`