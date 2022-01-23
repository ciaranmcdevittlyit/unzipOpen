/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"uzo/util"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		var err error
		var argument string

		argument = args[0]

		fileExists, err := util.FileExists(argument)
		if err != nil {
			fmt.Println(err.Error())
		}

		if fileExists {
			fileName, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("File %v does not exist", argument)
			return
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		util.Unzip(fileName, wd)

		os.Chdir(util.FilenameWithoutExtension(fileName))

		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		commandCode := exec.Command("code", wd)
		err = commandCode.Run()

		if err != nil {
			fmt.Println("VS Code executable file not found in %PATH%")
		}

	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}