package cmdTpl

const TplApi  =`package cmd

import (
	"github.com/moka-mrp/sword-core/kernel/server"
	"{{.ModuleName}}/app/api/router"
	"{{.ModuleName}}/config"
	"github.com/spf13/cobra"
)



//如果需要cobra捕获错误，则将错误返回即可，否则就自己处理就好了
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the api service of http",
	RunE: func(cmd *cobra.Command, args []string) error{
		//启动服务(启动服务中自带平滑重启机制额)
		err:=server.StartHttp(config.GetConf().Api,router.RegisterRoute)
		if err !=nil{
			return  err
		}
		return  nil
	},
}


func init() {
	rootCmd.AddCommand(apiCmd)
}`