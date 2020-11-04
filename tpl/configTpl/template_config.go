package configTpl

const TplConfig  =`package config

import (
	"github.com/moka-mrp/sword-core/config"
	"github.com/spf13/viper"
)

const (
	ProdEnv  = "production" //线上环境
	BetaEnv  = "beta"       //beta环境
	DevEnv   = "develop"    //开发环境
	LocalEnv = "local"      //本地环境
)

var (
	srvConf *Config
)


//------------------------配置文件解析
//@author sam@2020-08-07 09:30:41
//@todo 需要额外的配置请自己拓展即可
type Config struct {
	Env     string                ` + "`toml:\"Env\"`" + `
	ID     string                ` + "`toml:\"id\"`" + `
	Debug   bool                  ` + "`toml:\"Debug\"`" + `
	Log     config.LogConfig       ` + "`toml:\"Log\"`" + `
	Redis   config.RedisMultiConfig ` + "`toml:\"Redis\"`" + `
	Db      config.DbConfig        ` + "`toml:\"Db\"`" + `
	Api     config.ApiConfig       ` + "`toml:\"Api\"`" + `
	Etcd    config.EtcdConfig   ` + "`toml:\"Etcd\"`" + `
	Jwt config.JwtConfig ` + "`toml:\"Jwt\"`" + `
}

func newConfig() *Config {
	return new(Config)
}




//加载viper中的配置信息
//@author sam@2020-08-07 09:33:10
func Load() (*Config, error) {
	conf := newConfig()
	//直接反序列化为Struct
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}
	//默认值设置
	if len(conf.Log.Name) == 0 {
		conf.Log.Name = "sword"
	}
	srvConf = conf
	//--------------日志解析调试--------------------
	////全局
	//fmt.Printf("%+v\r\n",conf.Debug)
	//fmt.Printf("%+v\r\n",conf.Env)
	////日志
	//fmt.Printf("%+v\r\n",conf.Log)
	////Redis
	//fmt.Printf("%+v\r\n",conf.Redis)
	//fmt.Printf("%+v\r\n",conf.Redis["default"])
	//fmt.Printf("%+v\r\n",conf.Redis["center"])
	////Db
	//fmt.Printf("%+v\r\n",conf.Db.Option)
	//fmt.Printf("%+v\r\n",conf.Db.Master)
	////Api
	//fmt.Printf("%+v\r\n",conf.Api)
	//crontab
	//fmt.Printf("%+v\r\n",conf.Crontab)

	return conf, nil
}








//------------------------全局---------------------------------------------
//当前配置
//@author  sam@2020-08-07 09:37:40
func GetConf() *Config {
	return srvConf
}

//是否调试模式
//@author  sam@2020-08-07 09:37:31
func IsDebug() bool {
	return srvConf.Debug
}

//当前环境，默认本地开发
//@author  sam@2020-08-07 09:37:21
func GetEnv() string {
	if srvConf.Env == "" {
		return LocalEnv
	}
	return srvConf.Env
}

//是否当前环境
//@author sam@2020-08-07 09:37:10
func IsEnvEqual(env string) bool {
	return GetEnv() == env
}`