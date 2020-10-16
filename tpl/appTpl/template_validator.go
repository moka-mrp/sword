package appTpl

const TplValidator  =`package testFunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moka-mrp/xuanyuan/app/common/constants"
	. "github.com/moka-mrp/xuanyuan/app/common/controllers"
)


/*
 * validator.v9文档
 * 地址https://godoc.org/gopkg.in/go-playground/validator.v9
 * 列了几个大家可能会用到的，如有遗漏，请看上面文档
 */

//请求数据结构
type ValidatorRequest struct {
	//tips，因为组件required不管是没传值或者传 0 or "" 都通过不了，但是如果用指针类型，那么0就是0，而nil无法通过校验
	Id   *int64 `+"`json:\"id\" validate:\"required\" example:\"1\"`"+`
	Age  int `+"`json:\"age\" validate:\"required,gte=0,lte=130\" example:\"20\"`"+`
	Name *string `+"`json:\"name\" validate:\"required\" example:\"sam\"`"+`
	Email string `+"`json:\"email\" validate:\"required,email\" example:\"sam@qq.com\"`"+`
	Url  string `+"`json:\"url\" validate:\"required\" example:\"https://www.baidu.com/\"`"+`
	Mobile string `+"`json:\"mobile\" validate:\"required\" example:\"18565056618\"`"+`
	RangeNum int `+"`json:\"range_num\" validate:\"max=10,min=1\" example:\"3\"`"+` //[1,10]
	TestNum *int `+"`json:\"test_num\" validate:\"required,oneof=5 7 9\" example:\"7\"`"+` // 5|7|9
	Content *string `+"`json:\"content\" example:\"sam is a good man.\"`"+`
	Addresses []*Address `+"`json:\"addresses\" validate:\"required,dive,required\"  `"+`
}

type Address struct {
	Street string `+"`json:\"street\" validate:\"required\" example:\"新平街\"` "+`
	City   string `+"`json:\"city\" validate:\"required\" example:\"苏州\"`"+`
	Phone  string `+"`json:\"phone\" validate:\"required\" example:\"0518-86456342\"`"+`
}



func Validator(c *gin.Context) {
	request := new(ValidatorRequest)
	err := GenRequest(c, request)
	fmt.Println(err)
	if err != nil {
		Error(c, constants.CodeParamError)
		return
	}
	Success(c, request)
	return
}`