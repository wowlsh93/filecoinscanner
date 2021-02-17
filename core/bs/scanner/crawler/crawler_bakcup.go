/*
2021-02-10

Written by wowlsh93
*/

package crawler

//
//type CrawlerEth struct {
//	ethRpc   *eth.EthRPC
//	logging  *logrus.Logger
//	stopChan chan bool
//
//	startingBlock int
//	currentBlock  int
//	highestBlock  int
//}
//
//func NewEth(conf config.Configuration, stop chan bool, startingBlock int) CrawlerEth {
//
//	crawler := CrawlerEth{eth.New(conf.Scanner.Ethscanner.Node_listen_address),
//		flogging.GetLogger(),
//		stop, startingBlock, 0, 0}
//
//	return crawler
//}
//
//func (c *CrawlerEth) GetBlock() <-chan eth.BlockResult {
//
//	results := make(chan eth.BlockResult)
//
//	startBlock := c.startingBlock
//
//	go func() {
//		defer close(results)
//		for {
//
//			result := c.getChainData(startBlock)
//
//			select {
//			case <-c.stopChan:
//				return
//			case results <- result:
//				startBlock = startBlock + 1
//			}
//		}
//	}()
//
//	return results
//}
//
//func (c *CrawlerEth) getChainData(startblock int) eth.BlockResult {
//
//	var result eth.BlockResult
//	receivedBlock, err := c.ethRpc.EthGetBlockByNumber(startblock, true)
//	result = eth.BlockResult{err, receivedBlock}
//	return result
//
//}
//
//func (c *CrawlerEth) GetLastBlockHeight() int {
//
//	height, _ := c.ethRpc.EthLastBlockNumber()
//	return height
//}
