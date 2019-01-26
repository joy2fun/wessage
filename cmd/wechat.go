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
		wechat.RefreshToken()
		wechat.ListTemplates()
	},
}

var cmdSend = &cobra.Command{
	Use:   "send [content]",
	Short: "Send template message",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wechat.RefreshToken()
		wechat.SendTemplateMessage(args[0], receiver, templateID)
	},
}
