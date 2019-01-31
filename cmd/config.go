package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdConfig = &cobra.Command{
	Use:   "config",
	Short: "Configuration related commands",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var cmdConfigList = &cobra.Command{
	Use:   "list",
	Short: "Display configuration detail",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"key", "value"})
		fmt.Println("Loaded config file:", viper.ConfigFileUsed())

		keys := []string{
			"wx_appid",
			"wx_secret",
			"wx_template_id",
			"wx_receiver",
		}

		for _, v := range keys {
			table.Append([]string{
				v,
				viper.GetString(v),
			})
		}

		if ts := viper.GetInt64("wx_token_expired_at"); ts > 0 {
			table.Append([]string{
				"wx_token_expired_at",
				time.Unix(ts, 0).String(),
			})
		}

		table.Render()
	},
}

var cmdConfigSet = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Edit config settings",
	Long:  ``,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set(args[0], args[1])
		viper.WriteConfig()
	},
}
