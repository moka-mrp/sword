package appTpl


const TplCode=`package constants

//@link  https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
//todo 此处模仿 net/http/status中的写法，创建出属于我们自己的业务状态码额
//@author  sam@2020-08-24 14:41:40

const (
	//成功
	CodeSuccess = 0
	CodeSystemError = 500 	//系统错误 todo 一般是请求命中了panic
	CodeParamError = 40000  //参数错误
	//------------------------------------
	CodeRedisError=  40100  //redis服务器异常
	CodeMysqlError=  40101  //mysql服务器异常
	CodeMongodbError=  40102  //mongodb服务器异常
	CodeEtcdError=  40103  //Etcd服务器异常
    //------------------------------------
	CodeBase64Error = 40200  //base64解密错误
	CodeHmacHs256Error = 40201 //HmacHs256签名错误
	CodeRsaError = 40202
	CodeAesError = 40203
	//-----------------------------------------
	CodeCreateError=40300  //增加失败
	CodeUpdateError=40301  //修改失败
	CodeReadError=40302  //查询失败
	CodeDeleteError=40303  //删除失败
	CodeSaveError=40304  //保存失败

)





//状态码与错误信息的映射值
var codeText = map[int]string{
	CodeSuccess:     "ok",
	CodeSystemError: "the system had a bad cold",
	CodeParamError:  "invalid param",
	//------------------------------------
	CodeRedisError : "redis error",
	CodeMysqlError : "mysql error",
	CodeMongodbError : "mongodb error",
	CodeEtcdError : "etcd error",
	//------------------------------------
	CodeBase64Error:"invalid base64 literal",
	CodeHmacHs256Error :"invalid HmacHs256 signature",
	CodeRsaError : "rsa error",
	CodeAesError :"aes error",
	//-----------------------------------------
	CodeCreateError:"create failed",
	CodeUpdateError:"update failed",
	CodeReadError:"query failed",
	CodeDeleteError :"delete failed",
	CodeSaveError :"save failed",



}
//传递code码获取上述codeText定义的错误信息值
//@author sam@2020-04-07 10:47:32
func CodeText(code int) string {
	if msg, ok := codeText[code]; ok {
		return msg
	}
	return ""
}`