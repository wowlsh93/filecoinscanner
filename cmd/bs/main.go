/*
2021-02-10

Written by wowlsh93
*/

package main

import (
	"fmt"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner"
	"github.com/wowlsh93/filecoinscanner/core/bs/version"
	"github.com/spf13/cobra"
	"os"
)

var mainCmd = &cobra.Command{

	Use:   "bs",
	Short: "Sample scanner",
	Long:  `This application is simple scanner to learn filecoin`,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func main() {

	mainCmd.AddCommand(scanner.Cmd())
	mainCmd.AddCommand(version.Cmd())

	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
