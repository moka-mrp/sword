package cmd

import (
	"github.com/moka-mrp/sword/server"
	"github.com/spf13/cobra"
)

//todo go run main.go new   moye   -p /Tmp/  -m   github.com/moka-mrp/moye

//定义一个new 命令行
var newCmd = &cobra.Command{
	Use:   "new",
	Aliases: []string{"n"},
	Short: "create new project",
	RunE: func(cmd *cobra.Command, args []string)error {
		err:=server.RunNew(args)
		if err!=nil{
			return err
		}
		return  nil
	},
}


func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&server.P.Path,"path", "p","", "directory for create project, default: current position")
	newCmd.Flags().StringVarP(&server.P.ModuleName,"module", "m","", "project module name, for go mod init")
	//注意这里选项还没有解析到P中,Run中已经解析了
}



