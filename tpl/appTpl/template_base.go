package appTpl

const TplBase  = `package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/moka-mrp/sword-core/log/logger"
	"{{.ModuleName}}/app/common/utils"
	"{{.ModuleName}}/app/common/constants"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"
	"strings"
)
//-------------------接口统一的读取客户端传递的数据并自动解码成结构体且自动进行validate的验证-------------------------------------
/**
 * 将请求的json格式的body转换为request数据结构
 * @param c
 * @param request  传入request数据结构的指针 如 new(FastRequest)
 * @author sam@2020-04-07 11:25:08
 */
func GenRequest(c *gin.Context, request interface{}) (err error) {
	//读取客户端传递的body
	body, err := ReadBody(c)
	if err != nil {
		return
	}
	//json解码
	//json:= jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(body, request)
	if err == nil {
		validate := validator.New()
		errValidate := validate.Struct(request)
		if errValidate != nil {
			//记录验证错误的参数
			logger.Error(c.Request.URL.Path," [param validate err] ", errValidate)
			return errValidate
		}
	}else{
		logger.Error(c.Request.URL.Path," [param json unmarshal err] ",err)
		return  err
	}
	return nil
}

//如果客户端body传递的是["a","b","c"]
//@author sam@2020-09-05 09:22:47
func GenRequestSlice(c *gin.Context, request interface{}) (err error) {
	//读取客户端传递的body
	body, err := ReadBody(c)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, request)
	if err !=nil{
		logger.Error(c.Request.URL.Path," [param json unmarshal err] ",err)
		return  err
	}
	return nil
}



//重复读取body
//@reviser sam@2020-04-07 11:41:23
func ReadBody(c *gin.Context) (body []byte, err error) {
	// ReadAll 读取 r 中的所有数据，返回读取的数据和遇到的错误。
	//如果读取成功，则 err 返回 nil，而不是 EOF，因为 ReadAll 定义为读取所有数据，所以不会把 EOF 当做错误处理。
	body, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	//包装一个io.Reader，返回一个io.ReadCloser(读出来之后需要再归还额)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}
//----------------------接口统一的成功时返回方法------------------------------------------------------
/*
正常返回示例:
{
	"errcode": 0,
	"errmsg": "ok",
	"data": {
		"user_info": {
			"head_img_url": "http://wx.qlogo.cn/mmhead/iahdQicCC5VBRicqVNrrBMtAQJmF8sq6UgDG85y7DWreibgeWjMsJbLZQA/0",
			"nickname": "三省"
		},
		"group_id": "CBgAAoXi7WNS9_rVyw0y5",
		"channel_id": "CBgAAoXkpQY85ngSuUXnb",
		"channel_name": "秒射小枪手",
	  }
}
*/




/**
 * 接口统一的正确返回值
 *  Success(c, "sam is a good man.")
 *  Success(c,samStruct)
 *  Success(c,samMap)
 * @reviser sam@2020-04-07 10:30:16
 */
func Success(c *gin.Context, data interface{}) {

	if utils.IsNil(data){
		//data=make(map[string]string)
		data=make([]string,0)
	}
	c.JSON(http.StatusOK, gin.H{
		"errcode":     constants.CodeSuccess,
		"errmsg":         constants.CodeText(constants.CodeSuccess),
		"request_uri": c.Request.URL.Path,
		"data":        data,
	})
	c.Abort()
}
//--------------------------------- 接口统一的以及定制的失败时返回方法 ------------------------------------------


/*
错误返回示例:
{
    "errcode": -10000401,
    "errmsg": "Sam is sick"
}
*/

/**
 * 失败时返回方法,注意http状态码仍然是200，只是业务状态码code有所不同
 * Error(c,400, "路由不存在")
 * Error(c,500)
 * @todo 注意这里的code是业务代码,http code始终是200
 * @reviser sam@2020-04-07 10:32:17
 */
func Error(c *gin.Context, code int, msg ...string) {
	//获取错误信息
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = constants.CodeText(code)
	}

	c.JSON(http.StatusOK, gin.H{
		"errcode":        code,
		"errmsg":         message,
		"request_uri": c.Request.URL.Path,
	})
	c.Abort()
}

//路径不存在
//@todo 待引入HTML模板
//@author sam@2020-05-22 14:12:21
func Error404(c *gin.Context) {
	    //isAjax:=c.GetHeader("X-Requested-With")
		Error(c, http.StatusNotFound,http.StatusText(http.StatusNotFound))
}

func Error500(c *gin.Context) {
	Error(c,http.StatusInternalServerError,http.StatusText(http.StatusInternalServerError)) //服务器异常
}


//将传递的以,分隔开的检索参数切割成切片
//@author sam@2020-09-05 16:13:51
func GetStringArrayFromQuery(c *gin.Context,name, sep string) (arr []string) {
	val := strings.TrimSpace(c.Query(name))
	if len(val) == 0 {
		return
	}
	return strings.Split(val, sep)
}`