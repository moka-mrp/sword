package cmd

import (
	"fmt"
	"github.com/moka-mrp/sword/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var  versionCmd = &cobra.Command{
	Use:                        "version",
	Short:                      "sword version",
	Aliases:[]string{"v"},
	RunE: func(cmd *cobra.Command, args []string)error {
		fmt.Println(server.GetVersion())
		return nil
	},
}