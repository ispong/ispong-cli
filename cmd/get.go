/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "下拉博客代码 ispong-blogs",
	Long:  `下拉博客代码 ispong-blogs`,
	Run: func(cmd *cobra.Command, args []string) {

		// 判断token是否存在，否则让用户登录
		token := viper.GetString("token")
		if len(token) == 0 || token == "" {
			fmt.Println("请重新登录")
			os.Exit(1)
		}

		// 输入安装路径
		var blogPath string
		fmt.Print("请输入安装路径:")
		_, err := fmt.Scanln(&blogPath)
		if err != nil {
			return
		}

		// 判断文件夹不存在则新建
		_, err = os.Stat(blogPath)
		if os.IsNotExist(err) {
			err := os.Mkdir(blogPath, 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// 执行命令下载代码
		blogRepository := "https://github.com/ispong/ispong-blogs.git"
		gitCmd := exec.Command("git", "clone", blogRepository)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		gitCmd.Dir = blogPath
		err = gitCmd.Run()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		} else {
			fmt.Println("下载成功")
		}

		// 保存安装路径
		saveBlogPath(blogPath)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 保存配置
func saveBlogPath(blogPath string) {
	viper.Set("blog-path", blogPath)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
