package cmdTpl

const TplVersion = `package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

//输出sword版本
//@author sam@2020-08-11 10:05:03
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of sword",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sword version is v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}`