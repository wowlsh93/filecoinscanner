/*
2021-02-10

Written by wowlsh93
*/

package scanner

import (
	"fmt"
	"github.com/wowlsh93/filecoinscanner/common/flogging"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/bsmanager"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/config"
	"github.com/wowlsh93/filecoinscanner/core/bs/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logging *logrus.Logger

var detectMode string
var configPath string

func startCmd() *cobra.Command {

	scannerStartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is ./configurations/bs_config.yaml)")
	scannerStartCmd.PersistentFlags().StringVarP(&detectMode, "mode", "m", "", "mode - both, deposit, withdrawl (default is both")

	scannerStartCmd.MarkPersistentFlagRequired("mode")

	return scannerStartCmd
}

var scannerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start scanning..",
	Long:  `starts a scanning that interacts with the scanner node`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}

		cmd.SilenceUsage = true
		return serve(args)
	},
}

func serve(args []string) error {

	fmt.Println("Starting %s", version.GetInfo())
	var conf = config.InitConfig(configPath)

	// command line option has strong priority more than configuration file.
	conf.Server.Mode = detectMode
	flogging.InitLog(&conf)

	cleanUp := make(chan bool)
	bs := bsmanager.New(conf, cleanUp)
	bs.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//service lifecycle
	for {

		select {

		//case <- resume:
		//case <- suspend:
		//case <- reconfiguring:
		case <-cleanUp:
			fmt.Println("scanner cleanUP was Finished")
			time.Sleep(500 * time.Millisecond)
			return nil
		case sig := <-sigs:
			fmt.Println(sig)
			fmt.Println("scanner closing by signal")
			bs.Stop()
		}

		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("scanner exit")
	return nil
}
