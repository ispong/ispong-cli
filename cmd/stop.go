/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止服务，举例 i stop [blog|mysql]",
	Long:  `停止服务，举例 i stop [blog|mysql]`,
	Run: func(cmd *cobra.Command, args []string) {
		// i run blogs 启动博客
		if "blog" == args[0] {

			dockerCommand := "docker stop ispong-blogs"

			dockerCmd := exec.Command("bash", "-c", dockerCommand)
			dockerCmd.Stdout = os.Stdout
			dockerCmd.Stderr = os.Stderr
			err := dockerCmd.Run()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			} else {
				fmt.Println("停止成功")
			}
		}

		// i run mysql 启动本地mysql
		if "mysql" == args[0] {
			dockerCommand := "docker stop ispong-mysql"

			dockerCmd := exec.Command("bash", "-c", dockerCommand)
			dockerCmd.Stdout = os.Stdout
			dockerCmd.Stderr = os.Stderr
			err := dockerCmd.Run()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			} else {
				fmt.Println("停止成功")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
