package appTpl

const TplJwtTokenVerify  = `package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/moka-mrp/sword-core/http/ctxkit"
	"github.com/moka-mrp/sword-core/log/logger"
	"{{.ModuleName}}/app/common/controllers"
	"net/http"
)

// TokenVerify access token校验
//todo 我们在解析token的时候，具体token采用的是什么签名方式，alg中已经告知了
//@author sam@2020-08-18 09:53:48
func JwtTokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		//先判断是否存在
		tokenStr:=c.Request.Header.Get("Authorization")
		if len(tokenStr)==0{
			logger.GetLogger().Error("Authorization is not set")
			controllers.Error(c,401,"Unauthorized")
			return
		}
		//存在则做智能截取
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(ctxkit.GetJwtSecret(c)), nil
			})
		//多逻辑判断
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//获取token中保存的用户信息
			ctxkit.SetJwtClaims(c,claims)
			c.Next()
		} else {
			logger.GetLogger().Error(err)
			controllers.Error(c,http.StatusUnauthorized,http.StatusText(http.StatusUnauthorized))
			return
		}
	}
}`