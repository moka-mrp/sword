package bootstrapTpl

const TplBootstrap  =  `package bootstrap

import (
	"github.com/moka-mrp/sword-core/db"
	"github.com/moka-mrp/sword-core/etd"
	"github.com/moka-mrp/sword-core/kernel/close"
	"github.com/moka-mrp/sword-core/kernel/container"
	"github.com/moka-mrp/sword-core/log/logger"
	"github.com/moka-mrp/sword-core/rds"
	"{{.ModuleName}}/config"
)

//包全局变量
var App *container.Container

/**
 * 服务引导程序  第一个参数为注入别名，第二个参数为配置，第三个参数可选为是否懒加载
 * @todo 注意特殊资源都是惰性加载的,除非主动从容器中打捞才会真正创建，这样就可以统一资源启动了
 * @todo  缺少权重、优先级
 * @author  sam@2020-08-08 15:29:26
*/
func Bootstrap(conf *config.Config) (err error) {

	//容器
	App = container.App
	//注册日志类服务(非惰性)
	err = logger.Pr.Register(logger.SingletonMain, conf.Log)
	if err != nil {
		return
	}
	// 注册db服务(惰性的吧)
	err = db.Pr.Register(db.SingletonMain, conf.Db,true)
	if err != nil {
		return
	}
	//注册redis服务   注入别名(string) + 对应配置  + 是否惰性加载(false)
	err = rds.Pr.Register(rds.SingletonMain, conf.Redis,true)
	if err != nil {
		return
	}
	//注册etcd服务
	err=etd.Pr.Register(etd.SingletonMain,conf.Etcd,true)
	if err !=nil{
		return err
	}


	//注册应用停止时调用的关闭服务
	close.MultiRegister(db.Pr,rds.Pr,etd.Pr)



	return nil
}`