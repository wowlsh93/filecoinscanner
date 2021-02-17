/*
2021-02-10

Written by wowlsh93
*/

package bsmanager

//
//var logging *logrus.Logger
//
//// Managing Crawling (getting block ) and Analyzer (checking block )
//// with start/stop/resume/suspend/cleanup
//type BSManager struct {
//	crawler  crawler.Crawler
//	analyzer analyzer.Analyzer
//
//	stop    chan bool
//	resume  chan bool
//	suspend chan bool
//	cleanUp chan<- bool
//
//	conf config.Configuration
//}
//
//func New(conf config.Configuration, cleanup chan<- bool) *BSManager {
//
//	logging = flogging.GetLogger()
//
//	ret := new(BSManager)
//
//	ret.conf = conf
//	ret.stop = make(chan bool)
//	ret.cleanUp = cleanup
//
//	return ret
//}
//
//func (b *BSManager) Start() {
//	go b.run()
//}
//func (b *BSManager) Stop() {
//	b.stop <- true
//}
//
//func (b *BSManager) printRunningInfo() {
//	fmt.Println("===================================================")
//	logging.Info("BlockScanner running")
//	logging.Infof("BlockChain: %s", b.conf.Scanner.Ethscanner.Simbol)
//	logging.Infof("Node Address: %s", b.conf.Scanner.Ethscanner.Node_listen_address)
//	logging.Infof("Starting Block: %d", b.conf.Scanner.Ethscanner.Start_monitoring_block)
//	logging.Infof("Detecting Mode : %s", b.conf.Server.Mode)
//	fmt.Println("====================================================")
//}
//
//func (b *BSManager) run() {
//
//	b.printRunningInfo()
//
//	stopChan := make(chan bool)
//	blockChan := make(chan *eth.Block)
//
//	var analyzer_err error
//	b.analyzer, analyzer_err = analyzer.New(b.conf, blockChan, stopChan)
//
//	if analyzer_err != nil {
//		logging.Error(analyzer_err.Error())
//		return
//	}
//
//	b.analyzer.Start()
//	b.crawler = crawler.New(b.conf, stopChan, b.conf.Scanner.Ethscanner.Start_monitoring_block)
//
//	lastBlockNum := b.crawler.GetLastBlockHeight()
//	logging.Infof("Last Block Number is : %d", lastBlockNum)
//
//	///////////////////////////////////////////////////////////////////////
//	// main action of system
//	for result := range b.crawler.GetBlock() {
//		if result.Error != nil {
//			logging.Errorf("error: %v", result.Error)
//			continue
//		}
//
//		b.analyzer.BlockChan <- result.Block
//	}
//	///////////////////////////////////////////////////////////////////////
//
//	//block scanner lifecycle
//	for {
//		select {
//		//case <- resume:
//		//case <- suspend:
//		case <-b.stop:
//			stopChan <- true
//			return
//		}
//	}
//
//}
