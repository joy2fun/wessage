package cmd

import (
	"fmt"
	"os"

	"../wechat"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "wessage"}

var cmdTest = &cobra.Command{
	Use:   "test",
	Short: "test some func",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	initConfigCmd()
	rootCmd.AddCommand(cmdConfig)

	wechat.InitSendCmd()
	rootCmd.AddCommand(wechat.CmdTemplate)
	rootCmd.AddCommand(wechat.CmdSend)
	// rootCmd.AddCommand(cmdTest)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
