/*
2021-02-10

Written by wowlsh93
*/

package version

import (
	"fmt"
	"runtime"

	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/metadata"
	"github.com/spf13/cobra"
)

const ProgramName = "bs"

func Cmd() *cobra.Command {
	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "version",
	Short: "Print filecoinscanner blockscaner version.",
	Long:  `Print current version of the filecoinscanner blockscaner`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}

		cmd.SilenceUsage = true
		fmt.Print(GetInfo())
		return nil
	},
}

func GetInfo() string {

	return fmt.Sprintf("%s:\n Version: %s\n  Go version: %s\n"+
		" OS/Arch: %s\n",
		ProgramName, metadata.Version, runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))

}
