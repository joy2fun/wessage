package cmd

import (
	"../wechat"
	"github.com/spf13/cobra"
)

var cmdTemplate = &cobra.Command{
	Use:   "template",
	Short: "List all templates",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		wechat.ListTemplates()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		wechat.RefreshToken(refreshToken)
	},
}

var cmdSend = &cobra.Command{
	Use:   "send [content] [link]",
	Short: "Send template message",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			wechat.SendTemplateMessage(args[0], args[1], receiver, templateID)
		} else {
			wechat.SendTemplateMessage(args[0], "", receiver, templateID)
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		wechat.RefreshToken(refreshToken)
	},
}

var cmdUser = &cobra.Command{
	Use:   "user",
	Short: "List subscribed users",
	Long:  ``,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		wechat.ListUsers()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		wechat.RefreshToken(refreshToken)
	},
}
