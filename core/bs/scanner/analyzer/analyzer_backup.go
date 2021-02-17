/*
2021-02-10

Written by wowlsh93
*/

package analyzer

//
//var logging *logrus.Logger
//
//type Mode string
//
//const (
//	MODE_BOTH      = Mode("both")
//	MODE_DEPOSIT   = Mode("depoist")
//	MODE_WITHRAWAL = Mode("withrawal")
//)
//
//type Analyzer struct {
//	BlockChan   chan *eth.Block
//	reporter    reportprovider.ReportProvider
//	toReport    chan reportprovider.ReportForm
//	stopChan    chan bool
//	depositAddr *addressstorage.AddressDB
//	collectAddr *addressstorage.AddressSet
//	mode        Mode
//	blockNum    int
//}
//
//func New(conf config.Configuration, block chan *eth.Block, stop chan bool) (Analyzer, error) {
//
//	logging = flogging.GetLogger()
//
//	reportChan := make(chan reportprovider.ReportForm)
//	reporter := reportprovider.New(reportChan, conf.Scanner.Ethscanner.Notify_deposit_url,
//		conf.Scanner.Ethscanner.Notify_abnormal_withdrawal_url, stop)
//	reporter.Start()
//
//	var eMode Mode
//
//	switch conf.Server.Mode {
//	case "both", "b":
//		eMode = MODE_BOTH
//	case "depoist", "d":
//		eMode = MODE_DEPOSIT
//	case "withrawal", "w":
//		eMode = MODE_WITHRAWAL
//	default:
//		return Analyzer{}, errors.New("abnormal scanning mode is set!")
//	}
//
//	depositAddressdb, err := addressstorage.NewDB(conf.Scanner.Ethscanner.Deposit_account_db_path)
//
//	if err != nil {
//
//		return Analyzer{}, errors.New("addresssdb create fail !!")
//	}
//
//	analyzer := Analyzer{block, reporter, reportChan,
//		stop, depositAddressdb, addressstorage.NewSet(), eMode,
//		conf.Scanner.Ethscanner.Start_monitoring_block}
//
//	return analyzer, analyzer.loadingAddressList(conf.Scanner.Ethscanner.Deposit_account_list_path,
//		conf.Scanner.Ethscanner.Collect_account_list_path)
//}
//
//func (a *Analyzer) Start() {
//	go a.run()
//}
//
//func (a *Analyzer) run() {
//
//	for {
//		select {
//		case block := <-a.BlockChan:
//			a.analyze(block)
//		case <-a.stopChan:
//			return
//		}
//	}
//}
//
//func (a *Analyzer) loadingAddressList(depositAddressPath string, collectAddressPath string) error {
//
//	// deposit address setting
//	depfile, err := os.Open(depositAddressPath)
//	if err != nil {
//		return errors.New("deposit address list path is fail")
//	}
//	defer depfile.Close()
//
//	depscanner := bufio.NewScanner(depfile)
//	for depscanner.Scan() {
//		a.depositAddr.SetValue(depscanner.Text())
//	}
//
//	// collect address setting
//	colfile, err := os.Open(collectAddressPath)
//	if err != nil {
//		return errors.New("deposit address list path is fail")
//	}
//	defer colfile.Close()
//
//	colscanner := bufio.NewScanner(depfile)
//	for colscanner.Scan() {
//		a.collectAddr.SetValue(colscanner.Text())
//	}
//	return nil
//
//}
//
//func (a *Analyzer) analyze(block *eth.Block) {
//
//	a.blockNum = block.Number
//
//	switch a.mode {
//	case MODE_BOTH:
//		a.analyzeBoth(block)
//	case MODE_DEPOSIT:
//		a.analyzeReposit(block)
//	case MODE_WITHRAWAL:
//		a.analyzeWithrawal(block)
//	}
//}
//
//func (a *Analyzer) matchingError(err error) {
//	logging.Warningf("warning: %e\n", err)
//
//}
//func (a *Analyzer) matchingComplete() {
//	fmt.Printf("analysis is closed")
//}
//
//func (a *Analyzer) analyzeBoth(block *eth.Block) {
//	observable := rxgo.Just(block.Transactions)()
//	<-observable.ForEach(a.matchingStrategy, a.matchingError, a.matchingComplete)
//}
//
//func (a *Analyzer) analyzeReposit(block *eth.Block) {
//	observable := rxgo.Just(block.Transactions)()
//	<-observable.ForEach(a.matchingDepositStrategy, a.matchingError, a.matchingComplete)
//}
//
//func (a *Analyzer) analyzeWithrawal(block *eth.Block) {
//	observable := rxgo.Just(block.Transactions)()
//	<-observable.ForEach(a.matchingAbnWithrawalStrategy, a.matchingError, a.matchingComplete)
//}
//
//func (a *Analyzer) matchingStrategy(v interface{}) {
//	a.matchingDepositStrategy(v)
//	a.matchingAbnWithrawalStrategy(v)
//}
//
//func (a *Analyzer) matchingDepositStrategy(v interface{}) {
//	tx := v.(eth.Transaction)
//
//	// deposit check (someone sends to wowlsh93's deposit address)
//	if a.depositAddr.HasValue(tx.To) {
//		a.report(reportprovider.DEPOSIT_REPORT, tx.Hash, tx.From, tx.To)
//	}
//}
//
//func (a *Analyzer) matchingAbnWithrawalStrategy(v interface{}) {
//	tx := v.(eth.Transaction)
//
//	// abnormal withrawal check ( wowlsh93's managed deposit address -> non wowlsh93's managed address)
//	if a.depositAddr.HasValue(tx.From) && a.collectAddr.HasValue(tx.To) == false {
//		a.report(reportprovider.ABNWITHRWAL_REPORT, tx.Hash, tx.From, tx.To)
//	}
//}
//
//func (a *Analyzer) report(kind reportprovider.KIND, hash, from, to string) {
//
//	report := reportprovider.ReportForm{"eth", a.blockNum, hash, reportprovider.ABNWITHRWAL_REPORT, from, to}
//
//	a.toReport <- report
//
//}
