/*
2021-02-10

Written by wowlsh93
*/

package crawler

import (
	"errors"
	"fmt"
	"github.com/wowlsh93/filecoinscanner/common/flogging"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/config"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/crawler/fil"
	"github.com/sirupsen/logrus"
	"time"
)

type CrawlerFil struct {
	filRpc  *fil.FilRPC
	logging *logrus.Logger
	stop    chan bool

	dataChan chan fil.ChainDataResult

	startingBlock int
	currentBlock  int
	highestBlock  int
}

func NewFil(conf config.Configuration, startingBlock int, dataChan chan fil.ChainDataResult) CrawlerFil {

	crawler := CrawlerFil{fil.New(conf.Scanner.Filscanner.Node_listen_address),
		flogging.GetLogger(), make(chan bool), dataChan, startingBlock, 0, 0}

	return crawler
}

func (a *CrawlerFil) Start() {
	go a.runproxy(a.run)
}

func (a *CrawlerFil) Stop() {
	a.stop <- true
}

func (a *CrawlerFil) runproxy(f func()) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("crawler restart")
			go a.runproxy(f)

		}
	}()
	f()
}

func (c *CrawlerFil) run() {
	getnumber := c.startingBlock

	for {
		select {
		case <-time.After(100 * time.Millisecond):
			result, err := c.getChainData(getnumber)
			if err != nil {
				c.logging.Error(err.Error())
				continue
			}
			c.logging.Debugf("Crawling Ok  : %d", result.Height)
			c.dataChan <- result
			getnumber = getnumber + 1
		case <-c.stop:
			fmt.Println("CrawlerFil received stop signal")
			return
		}
	}
}

func (c *CrawlerFil) getChainData(getnumber int) (fil.ChainDataResult, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()

	tipset, err := c.filRpc.FilGetTipsetByHeight(getnumber)

	if err != nil {
		return fil.ChainDataResult{}, err
	}

	var allmsgs []fil.Messages
	for _, v := range tipset.Cids {
		msgs, _ := c.filRpc.FilGetMessagesByCID(v.Cid)
		allmsgs = append(allmsgs, msgs)
	}
	result := fil.ChainDataResult{err, getnumber, allmsgs}
	return result, err
}
