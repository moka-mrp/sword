package baseTpl

const TplBuild  =`#!/usr/bin/env bash
#sam@2020-08-11 11:02:54
#todo jenkins用来使用进行打包归档的(建议运维小哥做成Makefile)
#@todo 编译归档的时候只考虑版本，但不考虑部署的环境变量，部署的环境的变量是部署时候传递的
#@todo 检查归档效果: tar -ztvf ...


#自定义一个用来判断每次执行shell命令后是否存在错误，如果存在则直接以该错误码为结果值作为终结退出
#@author  sam@2020-08-11 11:05:27
function check_code() {
    EXCODE=$?
    if [ "$EXCODE" != "0" ]; then
        echo "build fail."
        exit $EXCODE
    fi
}

#设置环境变量
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN:$GOROOT/bin
export  GOPROXY=https://goproxy.cn
export  GO111MODULE=on
#设置默认值
CGO_ENABLED=0
GOARCH="amd64"
GOOS="$1"
if [ "${GOOS}" == "" ]; then
    GOOS="linux"
fi
VERSION="$2"
if [ "${VERSION}" == "" ]; then
    VERSION="v1.0.0"
fi
#项目目录的截取
CPATH=`+"`pwd`"+`
DIST=${CPATH}/dist
BIN=${CPATH}/bin
SOURCE=main.go
TARGET=shanbumin
#build之前打扫战场
rm -rf ${CPATH}/*.zip
rm -rf ${CPATH}/*.gz
rm -rf ${DIST}/*.gz
rm -rf ${BIN}/*.run

#编译函数
build(){
    echo "---------------build start: $GOOS $GOARCH--------------"
    TARGET_PACKAGE=${TARGET}-${VERSION}-${GOOS}-${GOARCH}
    #设置环境变量的同时进行编译
    env GOOS=$GOOS GOARCH=$GOARCH  CGO_ENABLED=$CGO_ENABLED  go build -o ${TARGET}  -v ${SOURCE}
    chmod +x ${TARGET}
    check_code
    mv ${TARGET}  ./bin
    check_code
    #打包归档
    if [ ${GOOS} == "windows" ];then
      zip ${TARGET_PACKAGE}.zip  ./bin/${TARGET}  deploy.sh  ./docs/conf  ./logs  ./resources  ./public
      mv ${TARGET_PACKAGE}.zip  ${DIST}/
    else
        tar --exclude=*.git --exclude=*.go  --exclude=*.log -czvf ${TARGET_PACKAGE}.tar.gz  \
         ./bin/${TARGET}  \
         ./deploy.sh \
         ./docs/conf \
         ./logs \
         ./resources \
         ./public
        #wrap
        TARGET_WRAP=${TARGET}.run
        cp /usr/local/bin/wrap.sh  ${TARGET_WRAP}
        check_code
        cat ${TARGET_PACKAGE}.tar.gz  >> ${TARGET_WRAP}
        chmod +x ${TARGET_WRAP}
        mv ${TARGET_PACKAGE}.tar.gz  ${DIST}/
        mv ${TARGET_WRAP}   ${BIN}/
    fi
    #删除编译的二进制吧
    rm -rf ./bin/${TARGET}
}
go mod tidy
build
`
