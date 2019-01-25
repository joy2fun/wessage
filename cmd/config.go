package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdConfig = &cobra.Command{
	Use:   "config",
	Short: "Configuration commands",
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
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"key", "value"})
		fmt.Println("Loaded config file:", home+"/.wessage.json")
		c := viper.AllSettings()
		for k := range c {
			if k != "wx_token" {
				var s []string
				if k == "wx_token_expired_at" {
					ts := viper.GetInt64(k)
					s = append(s, k, time.Unix(ts, 0).String())
				} else {
					s = append(s, k, viper.GetString(k))
				}
				table.Append(s)
			}
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

func InitConfig() {
	viper.SetConfigName(".wessage")
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 如果文件不存在，创建
	if _, err := os.Stat(home + "/.wessage.json"); os.IsNotExist(err) {
		s := []byte("{}")
		err := ioutil.WriteFile(home+"/.wessage.json", s, 0644)
		if err != nil {
			panic(err)
		}
	}

	viper.AddConfigPath(home)
	viper.SetConfigType("json")
	viper.ReadInConfig()
}

func initConfigCmd() {
	cmdConfig.AddCommand(cmdConfigList, cmdConfigSet)
}
