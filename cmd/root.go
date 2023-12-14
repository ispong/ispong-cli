/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ispong-cli",
	Short: "ispong个人工具集",
	Long:  `ispong个人工具集`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// 解析配置文件
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.i-cli/i-config.yml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {

	// 获取home目录
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 初始化配置文件信息
	viper.SetConfigFile(home + "/.i-cli/i-config.yml")

	// 判断配置文件是否存在
	if err := viper.ReadInConfig(); err != nil {

		// 判断文件夹是否存在，不存在则新建
		_, err := os.Stat(home + "/.i-cli")
		if os.IsNotExist(err) {
			err := os.Mkdir(home+"/.i-cli", 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// 判断文件是否存在，不存在则新建
		_, err = os.Stat(home + "/.i-cli/i-config.yml")
		if os.IsNotExist(err) {
			// 初始化配置
			viper.SetDefault("account", "ispong")
			// 持久化配置
			err = viper.SafeWriteConfigAs(home + "/.i-cli/i-config.yml")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
