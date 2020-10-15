package server

import (
	"bytes"
	"github.com/moka-mrp/sword/tpl"
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
	//初始化项目基本文件
	_tplTypeReadme = iota
	_tplTypeGitignore
	_tplTypeGoMod
	_tplTypeMain
	_tplTypeEnv
	_tplTypeEnvExample
	_tplTypeLogIgnore
	_tplTypeLogKeep
	_tplTypeBuildSh
	_tplTypeDeploySh
	_tplTypeDockerfile
	//
	_tplTypeCacheKey
	_tplTypeBannerListCache
	_tplTypeBannerListCacheTest
	_tplTypeConsoleKernel
	_tplTypeConsoleTest
	_tplTypeCommand
	_tplTypeConstantCommon
	_tplTypeConstantErrorCode
	_tplTypeConstantLogType
	_tplTypeControllerBase
	_tplTypeControllerTest
	_tplTypeEntity
	_tplTypeFormatter
	_tplTypeFormatterTest
	_tplTypeMiddleWare
	_tplTypeRoute
	_tplTypeJobBase
	_tplTypeJobKernel
	_tplTypeJobTest
	_tplTypeModel
	_tplTypeModelTest
	_tplTypeService
	_tplTypeUtil
	_tplTypeBootstrap
	_tplTypeConfig
	_tplTypeOption
	_tplTypeBuildBin
	_tplTypeBuildShell
	_tplTypeDocs
)


var (
	P project
	// 文件类型 => 文件路径
	files = map[int]string{
		//初始化项目基本文件 11
		_tplTypeReadme:     "/README.md",
		_tplTypeGitignore:  "/.gitignore",
		_tplTypeGoMod:      "/go.mod",
		_tplTypeMain:       "/main.go",
		_tplTypeEnv:        "/.env",
		_tplTypeEnvExample: "/.env.example",
		_tplTypeLogIgnore:        "/logs/.gitignore",
		_tplTypeLogKeep:        "/logs/.gitkeep",
		_tplTypeBuildSh:    "/build.sh",
		_tplTypeDeploySh:   "/deploy.sh",
		_tplTypeDockerfile: "/Dockerfile",
		////init caches
		//_tplTypeCacheKey:            "/app/caches/cache_key.go",
		//_tplTypeBannerListCache:     "/app/caches/bannerlistcache/banner_list.go",
		//_tplTypeBannerListCacheTest: "/app/caches/bannerlistcache/banner_list_test.go",
		////init console
		//_tplTypeConsoleKernel: "/app/console/kernel.go",
		//_tplTypeConsoleTest:   "/app/console/test.go",
		//_tplTypeCommand:       "/app/console/command.go",
		////init constant
		//_tplTypeConstantCommon:    "/app/constants/common/common.go",
		//_tplTypeConstantErrorCode: "/app/constants/errorcode/error_code.go",
		//_tplTypeConstantLogType:   "/app/constants/logtype/log_type.go",
		////init http
		//_tplTypeControllerBase: "/app/http/controllers/base.go",
		//_tplTypeControllerTest: "/app/http/controllers/test.go",
		//_tplTypeEntity:         "/app/http/entities/test.go",
		//_tplTypeFormatter:      "/app/http/formatters/bannerformatter/banner.go",
		//_tplTypeFormatterTest:  "/app/http/formatters/bannerformatter/banner_test.go",
		//_tplTypeMiddleWare:     "/app/http/middlewares/server_recovery.go",
		//_tplTypeRoute:          "/app/http/routes/route.go",
		////init job
		//_tplTypeJobBase:   "/app/jobs/basejob/base_job.go",
		//_tplTypeJobKernel: "/app/jobs/kernel.go",
		//_tplTypeJobTest:   "/app/jobs/test.go",
		////init model
		//_tplTypeModel:     "/app/models/bannermodel/banner.go",
		//_tplTypeModelTest: "/app/models/bannermodel/banner_test.go",
		////init service
		//_tplTypeService: "/app/services/bannerservice/banner.go",
		////init util
		//_tplTypeUtil: "/app/utils/.gitkeep",
		////init bootstrap
		//_tplTypeBootstrap: "/bootstrap/bootstrap.go",
		////init config
		//_tplTypeConfig: "/config/config.go",
		//_tplTypeOption: "/config/option.go",
		////init build
		//_tplTypeBuildBin:   "/build/bin/.gitignore",
		//_tplTypeBuildShell: "/build/shell/build.sh",
		////init docs
		//_tplTypeDocs: "/docs/docs.go",
	}

	//文件类型 =>文件模板内容
	//TODO: 映射不存在，则将会创建对应文件，但是内容为空而已
	tpls = map[int]string{
		//初始化项目基本文件
		_tplTypeReadme:     tpl.TplReadme,
		_tplTypeGitignore:  tpl.TplGitignore,
		_tplTypeGoMod:      tpl.TplGoMod,
		_tplTypeMain:       tpl.TplMain,
		_tplTypeEnv:        tpl.TplEnv,
		_tplTypeEnvExample: tpl.TplEnv,
		_tplTypeLogIgnore:  tpl.TplSimpleGitignore,
		_tplTypeBuildSh:    tpl.TplBuild,
		_tplTypeDeploySh:   tpl.TplDeploy,
		_tplTypeDockerfile: tpl.TplDockerfile,




		//_tplTypeCacheKey:            tpl._tplCacheKey,
		//_tplTypeBannerListCache:     tpl._tplBannerListCache,
		//_tplTypeBannerListCacheTest: tpl._tplBannerListCacheTest,
		//_tplTypeConsoleKernel:       tpl._tplConsoleKernel,
		//_tplTypeConsoleTest:         tpl._tplConsoleTest,
		//_tplTypeCommand:             tpl._tplCommand,
		//_tplTypeConstantCommon:      tpl._tplConstantCommon,
		//_tplTypeConstantErrorCode:   tpl._tplConstantErrorCode,
		//_tplTypeConstantLogType:     tpl._tplConstantLogType,
		//_tplTypeControllerBase:      tpl._tplControllerBase,
		//_tplTypeControllerTest:      tpl._tplControllerTest,
		//_tplTypeEntity:              tpl._tplEntity,
		//_tplTypeFormatter:           tpl._tplFormatter,
		//_tplTypeFormatterTest:       tpl._tplFormatterTest,
		//_tplTypeMiddleWare:          tpl._tplMiddleWare,
		//_tplTypeRoute:               tpl._tplRoute,
		//_tplTypeJobBase:             tpl._tplJobBase,
		//_tplTypeJobKernel:           tpl._tplJobKernel,
		//_tplTypeJobTest:             tpl._tplJobTest,
		//_tplTypeModel:               tpl._tplModel,
		//_tplTypeModelTest:           tpl._tplModelTest,
		//_tplTypeService:             tpl._tplService,
		//_tplTypeUtil:                tpl._tplUtil,
		//_tplTypeBootstrap:           tpl._tplBootstrap,
		//_tplTypeConfig:              tpl._tplConfig,
		//_tplTypeOption:              tpl._tplOption,
		//_tplTypeBuildBin:            tpl._tplBuildBin,
		//_tplTypeBuildShell:          tpl._tplBuildShell,
		//_tplTypeDocs:                tpl._tplDocs,
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
