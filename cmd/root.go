
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

//在没有任何子命令的情况下调用时，rootCmd表示基本命令
var rootCmd = &cobra.Command{
	Use:   "sword",
	Short: "sword tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Use "sword [command] -h,--help" for more information about a command.`)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


