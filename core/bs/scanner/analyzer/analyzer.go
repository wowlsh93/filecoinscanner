/*
2021-02-10

Written by wowlsh93
*/

package analyzer

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/wowlsh93/filecoinscanner/common/flogging"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/analyzer/addressstorage"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/analyzer/reportprovider"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/config"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/crawler/fil"
	"github.com/reactivex/rxgo/v2"
	"github.com/sirupsen/logrus"
	"os"
)

var logging *logrus.Logger

type Mode string

const (
	MODE_BOTH      = Mode("both")
	MODE_DEPOSIT   = Mode("depoist")
	MODE_WITHRAWAL = Mode("withrawal")
)

type Analyzer struct {
	DataChan    chan fil.ChainDataResult
	reporter    reportprovider.ReportProvider
	toReport    chan reportprovider.ReportForm
	stop        chan bool
	depositAddr *addressstorage.AddressDB
	collectAddr *addressstorage.AddressSet
	mode        Mode
	curHeight   int
}

func New(conf config.Configuration, datachan chan fil.ChainDataResult) (Analyzer, error) {

	logging = flogging.GetLogger()

	reportChan := make(chan reportprovider.ReportForm)
	reporter := reportprovider.New(reportChan, conf.Scanner.Ethscanner.Notify_deposit_url,
		conf.Scanner.Ethscanner.Notify_abnormal_withdrawal_url)
	reporter.Start()

	var eMode Mode

	switch conf.Server.Mode {
	case "both", "b":
		eMode = MODE_BOTH
	case "depoist", "d":
		eMode = MODE_DEPOSIT
	case "withrawal", "w":
		eMode = MODE_WITHRAWAL
	default:
		return Analyzer{}, errors.New("abnormal scanning mode is set!")
	}

	depositAddressdb, err := addressstorage.NewDB(conf.Scanner.Filscanner.Deposit_account_db_path)

	if err != nil {

		return Analyzer{}, errors.New("addresssdb create fail !!")
	}

	analyzer := Analyzer{datachan,
		reporter,
		reportChan,
		make(chan bool),
		depositAddressdb,
		addressstorage.NewSet(),
		eMode,
		0}

	return analyzer, analyzer.loadingAddressList(conf.Scanner.Filscanner.Deposit_account_list_path,
		conf.Scanner.Filscanner.Collect_account_list_path, conf.Scanner.Filscanner.Max_address_readbuffer_size)
}

func (a *Analyzer) Start() {
	go a.runproxy(a.run)
}

func (a *Analyzer) Stop() {
	a.stop <- true
}

func (a *Analyzer) runproxy(f func()) {
	defer func() {
		if v := recover(); v != nil {
			go a.runproxy(f)
		}
	}()

	f()
}

func (a *Analyzer) run() {

	for {
		select {
		case data := <-a.DataChan:
			if data.Error != nil {
				logging.Errorf("error: %v", data.Error)
				continue
			}
			a.analyze(data)
		case <-a.stop:
			fmt.Println("Analyzer received stop signal")
			return
		}
	}
}

func (a *Analyzer) loadingAddressList(depositAddressPath string, collectAddressPath string, maxbuffersize int) error {
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
	// deposit address setting
	depfile, err := os.Open(depositAddressPath)
	if err != nil {
		return errors.New("deposit address list path is fail")
	}
	defer depfile.Close()

	depscanner := bufio.NewScanner(depfile)

	buf := make([]byte, maxbuffersize)
	depscanner.Buffer(buf, maxbuffersize)
	for depscanner.Scan() {
		a.depositAddr.SetValue(depscanner.Text())
	}

	// collect address setting
	colfile, err := os.Open(collectAddressPath)
	if err != nil {
		return errors.New("deposit address list path is fail")
	}
	defer colfile.Close()

	colscanner := bufio.NewScanner(colfile)
	colscanner.Buffer(buf, maxbuffersize)
	for colscanner.Scan() {
		a.collectAddr.SetValue(colscanner.Text())
	}
	return nil

}

func (a *Analyzer) analyze(data fil.ChainDataResult) {

	a.curHeight = data.Height
	switch a.mode {
	case MODE_BOTH:
		a.analyzeBoth(data)
	case MODE_DEPOSIT:
		a.analyzeReposit(data)
	case MODE_WITHRAWAL:
		a.analyzeWithrawal(data)
	}
}

func (a *Analyzer) analyzeBoth(data fil.ChainDataResult) {

	observable := rxgo.Just(data.MessagesList)()
	<-observable.ForEach(a.matchingStrategy, a.matchingError, a.matchingComplete)
}

func (a *Analyzer) analyzeReposit(data fil.ChainDataResult) {
	observable := rxgo.Just(data.MessagesList)()
	<-observable.ForEach(a.matchingDepositStrategy, a.matchingError, a.matchingComplete)
}

func (a *Analyzer) analyzeWithrawal(data fil.ChainDataResult) {
	observable := rxgo.Just(data.MessagesList)()
	<-observable.ForEach(a.matchingAbnWithrawalStrategy, a.matchingError, a.matchingComplete)
}

func (a *Analyzer) matchingStrategy(v interface{}) {
	a.matchingDepositStrategy(v)
	a.matchingAbnWithrawalStrategy(v)
}

// deposit check (someone sends to wowlsh93's deposit address)
func (a *Analyzer) matchingDepositStrategy(v interface{}) {
	msgs := v.(fil.Messages)

	cidindex := 0
	for _, msg := range msgs.BlsMessages {

		if a.depositAddr.HasValue(msg.To) {
			a.report(reportprovider.DEPOSIT_REPORT, msgs.Cids[cidindex].Cid, msg.From, msg.To, msgs.Blockcid)
		}
		cidindex += cidindex
	}

	for _, msg := range msgs.SecpkMessages {
		if a.depositAddr.HasValue(msg.To) {
			a.report(reportprovider.DEPOSIT_REPORT, msgs.Cids[cidindex].Cid, msg.From, msg.To, msgs.Blockcid)
		}
		cidindex += cidindex
	}
}

// abnormal withrawal check ( wowlsh93's managed deposit address -> non wowlsh93's managed address)
func (a *Analyzer) matchingAbnWithrawalStrategy(v interface{}) {
	msgs := v.(fil.Messages)

	cidindex := 0
	for _, msg := range msgs.BlsMessages {

		if a.depositAddr.HasValue(msg.From) && a.collectAddr.HasValue(msg.To) == false {
			a.report(reportprovider.ABNWITHRWAL_REPORT, msgs.Cids[cidindex].Cid, msg.From, msg.To, msgs.Blockcid)
		}
		cidindex += cidindex
	}

	for _, msg := range msgs.SecpkMessages {
		if a.depositAddr.HasValue(msg.From) && a.collectAddr.HasValue(msg.To) == false {
			a.report(reportprovider.ABNWITHRWAL_REPORT, msgs.Cids[cidindex].Cid, msg.From, msg.To, msgs.Blockcid)
		}
		cidindex += cidindex
	}

}

func (a *Analyzer) matchingError(err error) {
	logging.Warningf("warning: %e\n", err)

}
func (a *Analyzer) matchingComplete() {
	logging.Debugf("analysis is finished : height: %d \n", a.curHeight)
}

func (a *Analyzer) report(kind reportprovider.KIND, hash, from, to, blockcid string) {

	report := reportprovider.ReportForm{"fil", a.curHeight, blockcid, hash, reportprovider.ABNWITHRWAL_REPORT, from, to}

	a.toReport <- report

}
