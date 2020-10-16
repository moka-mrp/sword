package cmdTpl

const TplOptions  =`package cmd

import (
	"github.com/spf13/pflag"
)


//启动服务时候设置的公共参数
//@author sam@2020-07-28 09:37:02
type SetupOptions struct {
	ConfigPath string
}
//添加选项
func (o *SetupOptions) AddFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(&o.ConfigPath, "config", "c", ".env", "config file (default is .env)")
}
//-------------------------------------------------------------------------------------
func newSetupOptions() *SetupOptions {
	return &SetupOptions{}
}`