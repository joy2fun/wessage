package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "wessage",
}

var cfgFile string
var receiver string
var templateID string

// Init 初始化
func Init() {
	cobra.OnInitialize(initConfig)

	// 配置相关命令
	rootCmd.AddCommand(cmdConfig)
	cmdConfig.AddCommand(cmdConfigList, cmdConfigSet)

	// 微信相关命令
	cmdSend.Flags().StringVarP(&templateID, "template", "t", "", "template id. use 'wx_template_id' in config if not present")
	cmdSend.Flags().StringVarP(&receiver, "receiver", "r", "", "receiver's openId. use 'wx_receiver' in config if not present")
	rootCmd.AddCommand(cmdSend)
	rootCmd.AddCommand(cmdTemplate)

	// 全局选项，自定义配置文件的路径
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.wessage.json)")
}

// Execute 初始化
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
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
		viper.SetConfigName(".wessage")
		viper.SetConfigType("json")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

}
