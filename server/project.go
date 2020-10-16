package server

import (
	"bytes"
	"github.com/moka-mrp/sword/tpl/appTpl"
	"github.com/moka-mrp/sword/tpl/baseTpl"
	"github.com/moka-mrp/sword/tpl/bootstrapTpl"
	"github.com/moka-mrp/sword/tpl/cmdTpl"
	"github.com/moka-mrp/sword/tpl/configTpl"
	"github.com/moka-mrp/sword/tpl/publicTpl"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)


//项目配置
type project struct {
	Name       string //项目名，如blog
	Path       string
	ModuleName string // 支持项目的自定义module名 （go.mod init）
}

//将所有要拷贝的文件都定义成常量类型
const (
	//初始化项目基本文件 9
	_tplTypeEnv =iota
	_tplTypeEnvExample
	_tplTypeGitignore
	_tplTypeBuildSh
	_tplTypeDeploySh
	_tplTypeDockerfile
	_tplTypeGoMod
	_tplTypeMain
	_tplTypeReadme
	//初始化app目录
	_tplTypeApp01
	_tplTypeApp02
	_tplTypeApp03
	_tplTypeApp04
	_tplTypeApp05
	_tplTypeApp06
	_tplTypeApp07
	_tplTypeApp08
	_tplTypeApp09
	_tplTypeApp10
	_tplTypeApp11
	_tplTypeApp12
	_tplTypeApp13
	_tplTypeApp14
	_tplTypeApp15
	_tplTypeApp16
    //初始化bin目录
	_tplTypeBin01
    _tplTypeBin02
	//初始化bootstrap目录
	_tplTypeBootstrap01
	_tplTypeBootstrap02
	//初始化cmd目录
	_tplTypeCmd01
	_tplTypeCmd02
	_tplTypeCmd03
	_tplTypeCmd04
	//初始化config目录
	_tplTypeConfig01
	//初始化dist目录
	_tplTypeDist01
	_tplTypeDist02
	//初始化docs目录
	_tplTypeDocs01
	//初始化logs目录
	_tplTypeLogs01
	_tplTypeLogs02
	//初始化pkg目录
	_tplTypePkg01
	//初始化providers目录
	_tplTypeProviders01
	//初始化public目录
	_tplTypePublic01
	_tplTypePublic02
	_tplTypePublic03
	_tplTypePublic04
	//初始化resources目录
	_tplTypeResources01
	_tplTypeResources02
	_tplTypeResources03
)

var (
	P project
	// 文件类型 => 文件路径
	files = map[int]string{
		//初始化项目基本文件 9
		_tplTypeReadme:     "/README.md",
		_tplTypeGitignore:  "/.gitignore",
		_tplTypeGoMod:      "/go.mod",
		_tplTypeMain:       "/main.go",
		_tplTypeEnv:        "/.env",
		_tplTypeEnvExample: "/.env.example",
		_tplTypeBuildSh:    "/build.sh",
		_tplTypeDeploySh:   "/deploy.sh",
		_tplTypeDockerfile: "/Dockerfile",
		//初始化app目录  TODO: tree app
		_tplTypeApp01:     "/app/api/controller/adminFunc/add_account.go",
		_tplTypeApp02:     "/app/api/controller/adminFunc/login.go",
		_tplTypeApp03:     "/app/api/controller/healthFunc/ping.go",
		_tplTypeApp04:     "/app/api/controller/testFunc/fast.go",
		_tplTypeApp05:     "/app/api/controller/testFunc/log.go",
		_tplTypeApp06:     "/app/api/controller/testFunc/redis.go",
		_tplTypeApp07:     "/app/api/controller/testFunc/validator.go",
		_tplTypeApp08:     "/app/api/router/router.go",
		_tplTypeApp09:     "/app/common/constants/code.go",
		_tplTypeApp10:     "/app/common/controllers/base.go",
		_tplTypeApp11:     "/app/common/middlewares/ctxkit.go",
		_tplTypeApp12:     "/app/common/middlewares/jwt_token_verify.go",
		_tplTypeApp13:     "/app/common/middlewares/server_recovery.go",
		_tplTypeApp14:     "/app/common/utils/common.go",
		_tplTypeApp15:     "/app/common/utils/common_test.go",
		//初始化bin目录
		_tplTypeBin01:    "/bin/.gitignore",
		_tplTypeBin02:    "/bin/.gitkeep",
		//初始化bootstrap目录
		_tplTypeBootstrap01: "/bootstrap/bootstrap.go",
		_tplTypeBootstrap02:"/bootstrap/welcome.go",
		//初始化cmd目录
		_tplTypeCmd01: "/cmd/api.go",
		_tplTypeCmd02: "/cmd/options.go",
		_tplTypeCmd03: "/cmd/root.go",
		_tplTypeCmd04: "/cmd/version.go",
		//初始化config目录
		_tplTypeConfig01: "/config/config.go",
		//初始化dist目录
		_tplTypeDist01: "/dist/.gitignore",
		_tplTypeDist02: "/dist/.gitkeep",
		//初始化docs目录
		_tplTypeDocs01: "/docs/.gitkeep",
		//初始化logs目录
		_tplTypeLogs01: "/logs/.gitignore",
		_tplTypeLogs02: "/logs/.gitkeep",
		//初始化pkg目录
		_tplTypePkg01: "/pkg/.gitkeep",
		//初始化providers目录
		_tplTypeProviders01: "/providers/.gitkeep",
		//初始化public目录
		_tplTypePublic01: "/public/assets/admin/index.html",
		_tplTypePublic02: "/public/assets/admin/css/.gitkeep",
		_tplTypePublic03: "/public/assets/admin/img/.gitkeep",
		_tplTypePublic04: "/public/assets/admin/js/.gitkeep",
		//初始化resources目录
		_tplTypeResources01: "/resources/cipher/.gitkeep",
		_tplTypeResources02: "/resources/db/.gitkeep",
		_tplTypeResources03: "/resources/views/admin/.gitkeep",

	}

	//文件类型 =>文件模板内容
	//TODO: 映射不存在，则将会创建对应文件，但是内容为空而已
	tpls = map[int]string{
		//初始化项目基本文件  9
		_tplTypeEnv :baseTpl.TplEnv,
		_tplTypeEnvExample:baseTpl.TplEnv,
		_tplTypeGitignore : baseTpl.TplGitignore,
		_tplTypeBuildSh  : baseTpl.TplBuild,
		_tplTypeDeploySh  : baseTpl.TplDeploy,
		_tplTypeDockerfile : baseTpl.TplDockerfile,
		_tplTypeGoMod  : baseTpl.TplGoMod,
		_tplTypeMain  : baseTpl.TplMain,
		_tplTypeReadme : baseTpl.TplReadme,

		//初始化app目录  TODO: tree app
		_tplTypeApp01:     appTpl.TplAddAccount,
		_tplTypeApp02:     appTpl.TplLogin,
		_tplTypeApp03:     appTpl.TplPing,
		_tplTypeApp04:     appTpl.TplFast,
		_tplTypeApp05:     appTpl.TplLog,
		_tplTypeApp06:     appTpl.TplRedis,
		_tplTypeApp07:     appTpl.TplValidator,
		_tplTypeApp08:     appTpl.TplRouter,
		_tplTypeApp09:     appTpl.TplCode,
		_tplTypeApp10:     appTpl.TplBase,
		_tplTypeApp11:     appTpl.TplCtxkit,
		_tplTypeApp12:     appTpl.TplJwtTokenVerify,
		_tplTypeApp13:     appTpl.TplServerRecovery,
		_tplTypeApp14:     appTpl.TplCommon,
		_tplTypeApp15:     appTpl.TplCommonTest,
		//初始化bin目录
		_tplTypeBin01:     baseTpl.TplSimpleGitignore,
		//_tplTypeBin02:    "/bin/.gitkeep",
		//初始化bootstrap目录
		_tplTypeBootstrap01: bootstrapTpl.TplBootstrap,
		_tplTypeBootstrap02: bootstrapTpl.TplWelcome,
		//初始化cmd目录
		_tplTypeCmd01: cmdTpl.TplApi,
		_tplTypeCmd02: cmdTpl.TplOptions,
		_tplTypeCmd03: cmdTpl.TplRoot,
		_tplTypeCmd04: cmdTpl.TplVersion,
		//初始化config目录
		_tplTypeConfig01: configTpl.TplConfig,
		//初始化dist目录
		_tplTypeDist01: baseTpl.TplSimpleGitignore,
		//_tplTypeDist02: "/dist/.gitkeep",
		//初始化docs目录
		//_tplTypeDocs01: "/docs/.gitkeep",
		//初始化logs目录
		_tplTypeLogs01: baseTpl.TplSimpleGitignore,
		//_tplTypeLogs02: "/logs/.gitkeep",
		//初始化pkg目录
		//_tplTypePkg01: "/pkg/.gitkeep",
		//初始化providers目录
		//_tplTypeProviders01: "/providers/.gitkeep",
		//初始化public目录
		_tplTypePublic01: publicTpl.TplIndex,
		//_tplTypePublic02: "/public/assets/admin/css/.gitkeep",
		//_tplTypePublic03: "/public/assets/admin/img/.gitkeep",
		//_tplTypePublic04: "/public/assets/admin/js/.gitkeep",
		//初始化resources目录
		//_tplTypeResources01: "/resources/cipher/.gitkeep",
		//_tplTypeResources02: "/resources/db/.gitkeep",
		//_tplTypeResources03: "/resources/views/admin/.gitkeep",

	}
)

//创建项目
//@author  sam@2020-10-14 13:40:45
func create() (err error) {
	//创建要生成的项目全路径
	if err = os.MkdirAll(P.Path, 0755); err != nil {
		return
	}
	//遍历要拷贝的所有的文件
	for t, v := range files {
		//如果拷贝的文件是多层的，则先创建其所在目录路径
		i := strings.LastIndex(v, "/")
		if i > 0 {
			dir := v[:i]
			if err = os.MkdirAll(P.Path+dir, 0755); err != nil {
				return
			}
		}
		//将对应模板文件的内容变量替换成功之后，然后按照对应的路径进行拷贝
		if err = write(P.Path+v, tpls[t], P); err != nil {
			return
		}
	}
	return
}

//创建文件
//@author sam@2020-10-14 14:20:23
func write(name, tpl string, data interface{}) (err error) {
	//解析对应模板
	body, err := parse(tpl, data)
	if err != nil {
		return
	}
	//创建文件
	return ioutil.WriteFile(name, body, 0644)
}

//解析模板并获取对应的解析后的模板内容
//@author  sam@2020-10-14 15:10:37
func parse(s string, data interface{}) ([]byte, error) {
	t, err := template.New("").Parse(s)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
