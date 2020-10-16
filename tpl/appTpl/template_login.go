package appTpl

const TplLogin  =`package adminFunc

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/moka-mrp/sword-core/http/ctxkit"
	"{{.ModuleName}}/app/common/constants"
	"{{.ModuleName}}/app/common/controllers"
	"{{.ModuleName}}/config"
	"time"
)

type LoginRequest struct {
	Email  string `+"`json:\"email\" validate:\"required\"`"+`
	Password string `+"`form:\"password\" validate:\"required\"`"+`
}

//提交登录
//todo 这里仅仅是个展示额
//curl -XPOST http://localhost:9999/v1/admin/login -H "Content-Type: application/json" -d  '{"email":"sam@qq.com","password":"123456"}'
//@author sam@2020-06-15 08:24:50
func Login(c *gin.Context) {

	//1.get request
	request := new(LoginRequest)
	err := controllers.GenRequest(c, request)
	if err != nil {
		controllers.Error(c,constants.CodeParamError)
		return
	}
	//2.从数据库中获取用户(可以从服务层获取)
	//略
	//3.验证密码是否正确
	//略
	//4.验证成功之后，返回访问令牌
	jwtSecret:=ctxkit.GetJwtSecret(c)
	exp:=time.Now().Add(time.Hour * time.Duration(config.GetConf().Jwt.Exp)).Unix()
	//创建一个新的令牌对象，指定签名方法和您希望包含的声明。
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"exp":exp, // jwt的过期时间，这个过期时间必须要大于签发时间
		"name":"sam",
		"email":"shanbumin@qq.com",
	})
	//生成token串
	tokenString, err := token.SignedString([]byte(jwtSecret))
	//5.返回
	response:=make(map[string]interface{},2)
	response["token_type"]="Bearer"
	response["access_token"]=tokenString
	response["expires_in"]=exp

	controllers.Success(c,response)

}`