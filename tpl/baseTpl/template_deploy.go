package baseTpl

const TplDeploy  = `#!/usr/bin/env bash

# Author   Sam  <shanbumin@qq.com>
# Date     2020-08-11 11:44:22
# bash deploy.sh  start
#$1是指明开启指令,如 start|stop|reload|restart
#$@所有参数，不包含命令本身
#todo 真机部署这里后续可以考虑换成systemctl后台托管程序
#todo 这里是部署脚本，运行环境肯定在这里啊

echo "PID for deploy.sh = $$"
tell(){
    echo "Usage: $0 run|start|stop|rerun|restart";
    echo "      eg: bash deploy.sh run";
    exit 404;
}

#自定义一个用来判断每次执行shell命令后是否存在错误，如果存在则直接以该错误码为结果值作为终结退出
function check_code() {
    EXCODE=$?
    if [ "$EXCODE" != "0" ]; then
        echo "deploy fail."
        tell
        exit $EXCODE
    fi
}



#(1)参数设置默认值
#export ENV=local
do="$1"
#penv="$2"
env="$ENV"
EXEC_NAME="shanbumin"
EXEC="./bin/${EXEC_NAME}"
#去掉deploy.sh必须的1个参数，后续参数都传给可执行的go程序  比如 ./shanbumin.run  run api ==> go run main.go api
shift 1
#(2)判断env是否指定
#参数指明的配置文件优先级高于环境变量传递的额
#todo 后续转成配置中心化之后就可以舍弃这一步骤了
#if [ ${penv} ];then
#  env="${penv}"
#fi

if [ ${env} ];then
  cp  -f ./docs/conf/${env}.env   ./.env
  check_code
else
   tell
fi

#-----------------------------------------------------
#前台运行，方便docker部署
run(){
    echo running... ${EXEC} ${env}
    ${EXEC} $@
    check_code
}
#后台运行,启动方法
start(){
    echo "starting ${env}"
    #nohup ${EXEC}  > ./nohup.log 2>&1 &
    nohup ${EXEC} $@ > /dev/null 2>&1 &
    echo "started ${env}"
}

#后台、前台运行的停止方法
stop(){
    PIDS=`+"`ps -ef|grep \"${EXEC}\"|grep -v 'grep'|awk '{print $2}'`"+`
    echo "$PIDS"
    for pid in $PIDS
    do
        echo kill $pid
        kill -15 $pid #优雅杀死，方便go程序接收信号SIGTERM，做结束工作
    done
}
#后台重启
restart(){
    stop && start $@
}
#前台重启
rerun(){
    stop && run $@
}
#----------------------------------------------------------------------------------------------------------

#(3)根据do调用对应方法
case "${do}" in
    run)
        run "$@";;
    start)
        start "$@";;
    stop)
        stop "$@";;
    rerun)
        rerun "$@";;
    restart)
        restart "$@";;
    *)
       tell
esac

exit 0`