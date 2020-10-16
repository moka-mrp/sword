package cmdTpl

const TplRoot  =`package cmd

import (
	"errors"
	"fmt"
	"github.com/moka-mrp/sword-core/kernel/server"
	"{{.ModuleName}}/bootstrap"
	"{{.ModuleName}}/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	opt  *SetupOptions
	errConfigPathRequired = errors.New("The config file path is required")
)




//@author sam@2020-08-06 13:42:33
//项目启动在没有指明任何命令的情况下，rootCmd作为默认执行的命令
var rootCmd = &cobra.Command{
	Use:   "sword", //励志打造一把举世无双的剑(一剑独尊)
	Short: "Sword is a very fast service generator.",
	Run: func(cmd *cobra.Command, args []string) {
		//注意我们目前没有使用到根命令，所以这里提醒即可。
		fmt.Println(`+"`Use \"sword [command] -h,--help\" for more information about a command.`"+`)
	},
}

//Execute将所有子命令添加到root命令并适当设置命令选项
//这由main.main（）调用。 它只需要对rootCmd发生一次。
//@author sam@2020-08-06 13:44:02
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

//初始化
//author sam@2020-08-06 14:03:31
func init() {
	cobra.OnInitialize(initConfig)
	//配置命令行选项
	opt = newSetupOptions()
	opt.AddFlags(rootCmd.PersistentFlags())
}

//OnInitialize将传递的函数initConfig设置为在每个命令调用Execute方法回调函数之前执行
//@todo  这里后续升级成走配置中心化，支持环境变量、命令行传参、配置文件等多形式
//@todo 优先级从高到低  环境变量-配置文件-配置中心(配置中心先暂时不打通...)
//export MYENV=black
//export -p |grep MYENV
//unset   MYENV
//@author  sam@2020-08-06 14:09:40
func initConfig() {
	//1.欢迎光临
	fmt.Println(bootstrap.AppWelcome)
	//2.设置viper相关属性
	if opt.ConfigPath == ""{
		log.Fatal(errConfigPathRequired)
	}
	viper.SetConfigFile(opt.ConfigPath)
	viper.SetConfigType("toml")
	viper.AutomaticEnv() 	//环境变量
	//3.解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read in config err:%s",err.Error())
	}
	fmt.Println("[INIT-conf] using config file:", viper.ConfigFileUsed(),"\r\n")
	//4.加载公共配置
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("load common config err:%s",err.Error())
	}
	server.SetDebug(conf.Debug)
	//5.启动引导程序
	err= bootstrap.Bootstrap(conf)
	if err != nil {
		log.Fatalf("bootstrap err:%s",err.Error())
	}
}`