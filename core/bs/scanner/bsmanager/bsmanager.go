/*
2021-02-10

Written by wowlsh93
*/

package bsmanager

import (
	"fmt"
	"github.com/wowlsh93/filecoinscanner/common/flogging"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/analyzer"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/config"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/crawler"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/crawler/fil"
	"github.com/sirupsen/logrus"
	"time"
)

var logging *logrus.Logger

// Managing Crawling (getting block ) and Analyzer (checking block )
// with start/stop/resume/suspend/cleanup
type BSManager struct {
	crawler  crawler.CrawlerFil
	analyzer analyzer.Analyzer

	stop    chan bool
	resume  chan bool
	suspend chan bool

	conf config.Configuration

	cleanUp chan bool
}

func New(conf config.Configuration, cleanup chan bool) *BSManager {

	logging = flogging.GetLogger()

	ret := new(BSManager)

	ret.conf = conf
	ret.stop = make(chan bool)
	ret.cleanUp = cleanup

	return ret
}

func (b *BSManager) Start() {
	go b.run()
}
func (b *BSManager) Stop() {
	b.stop <- true
}

func (b *BSManager) printRunningInfo() {
	fmt.Println("===================================================")
	logging.Info("BlockScanner running")
	logging.Infof("BlockChain: %s", b.conf.Scanner.Ethscanner.Simbol)
	logging.Infof("Node Address: %s", b.conf.Scanner.Ethscanner.Node_listen_address)
	logging.Infof("Starting Block: %d", b.conf.Scanner.Ethscanner.Start_monitoring_block)
	logging.Infof("Detecting Mode : %s", b.conf.Server.Mode)
	fmt.Println("====================================================")
}

func (b *BSManager) run() {

	b.printRunningInfo()
	dataChan := make(chan fil.ChainDataResult)
	// Analyzer init
	var analyzer_err error
	b.analyzer, analyzer_err = analyzer.New(b.conf, dataChan)

	if analyzer_err != nil {
		logging.Error(analyzer_err.Error())
		return
	}
	b.analyzer.Start()

	// Crawler init
	b.crawler = crawler.NewFil(b.conf, b.conf.Scanner.Filscanner.Start_monitoring_block, dataChan)
	b.crawler.Start()

	//block scanner lifecycle
	for {
		select {
		case <-b.stop:
			fmt.Println("bsmanager received stop signal")

			b.analyzer.Stop()
			b.crawler.Stop()
			b.cleanUp <- true
		//case <- resume:
		//case <- suspend:
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

}
