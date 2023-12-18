/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "docker 启动本地服务",
	Long:  `docker 启动启动本地服务`,
	Run: func(cmd *cobra.Command, args []string) {

		// i run blogs 启动博客
		if "blog" == args[0] {

			// 博客地址
			blogDirPath := viper.GetString("blog-path") + "/ispong-blogs"

			dockerCommand := "docker run " +
				"--name ispong-blogs " +
				"-v" + blogDirPath + ":/hexo " +
				"-p 4000:4000 " +
				"-d isxcode/hexo"

			dockerCmd := exec.Command("bash", "-c", dockerCommand)
			dockerCmd.Stdout = os.Stdout
			dockerCmd.Stderr = os.Stderr
			err := dockerCmd.Run()
			if err != nil {
				dockerCommand := "docker start ispong-blogs"
				dockerCmd := exec.Command("bash", "-c", dockerCommand)
				dockerCmd.Stdout = os.Stdout
				dockerCmd.Stderr = os.Stderr
				err = dockerCmd.Run()
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				} else {
					fmt.Println("启动成功")
					fmt.Println("访问地址：http://localhost:4000")
				}
			} else {
				fmt.Println("启动成功")
				fmt.Println("访问地址：http://localhost:4000")
			}
		}

		// i run mysql 启动本地mysql
		if "mysql" == args[0] {

			// 数据地址
			dataDirPath := "/Users/ispong/data/mysql/data"

			// 配置地址
			confDirPath := "/Users/ispong/data/mysql/conf.d"

			// 目录不存在则新建
			_, err := os.Stat(dataDirPath)
			if os.IsNotExist(err) {
				err := os.Mkdir(dataDirPath, 0755)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}

			// 目录不存在则新建
			_, err = os.Stat(confDirPath)
			if os.IsNotExist(err) {
				err := os.Mkdir(confDirPath, 0755)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}

			// docker命令
			dockerCommand := "docker run " +
				"--name ispong-mysql " +
				"--privileged=true " +
				"-p 30306:3306 " +
				"-e MYSQL_ROOT_PASSWORD=ispong123 " +
				"-e MYSQL_DATABASE=ispong_db " +
				"-v " + dataDirPath + ":/var/lib/mysql " +
				"-v " + confDirPath + ":/etc/mysql/conf.d " +
				"-d mysql:8.0"

			// 执行命令
			dockerCmd := exec.Command("bash", "-c", dockerCommand)
			dockerCmd.Stdout = os.Stdout
			dockerCmd.Stderr = os.Stderr
			err = dockerCmd.Run()
			if err != nil {
				dockerCommand := "docker start ispong-mysql"
				dockerCmd := exec.Command("bash", "-c", dockerCommand)
				dockerCmd.Stdout = os.Stdout
				dockerCmd.Stderr = os.Stderr
				err = dockerCmd.Run()
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				} else {
					fmt.Println("启动成功")
					fmt.Println("url：jdbc:mysql://localhost:30306/ispong_db")
					fmt.Println("username：root")
					fmt.Println("password：ispong123")
				}
			} else {
				fmt.Println("启动成功")
				fmt.Println("url：jdbc:mysql://localhost:30306/ispong_db")
				fmt.Println("username：root")
				fmt.Println("password：ispong123")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
