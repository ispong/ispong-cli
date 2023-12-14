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

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "提交博客代码",
	Long:  `提交博客代码`,
	Run: func(cmd *cobra.Command, args []string) {
		executeCommand := "cd " + viper.GetString("blog-path") + "/ispong-blogs && git add . && git commit -m \":memo: 写博客\" && git push origin main"
		result := exec.Command("bash", "-c", executeCommand)
		result.Stdout = os.Stdout
		result.Stderr = os.Stderr

		err := result.Run()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("提交代码成功")
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
