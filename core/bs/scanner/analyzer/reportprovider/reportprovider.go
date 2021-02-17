package reportprovider

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wowlsh93/filecoinscanner/common/flogging"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ReportProvider struct {
	report     chan ReportForm
	depositUrl string
	abnWUrl    string

	stop chan bool

	logging *logrus.Logger
}

func New(report chan ReportForm, depoistNotiUrl, abnWNotiUrl string) ReportProvider {

	reporter := ReportProvider{report, depoistNotiUrl, abnWNotiUrl, make(chan bool), flogging.GetLogger()}
	return reporter
}

func (r *ReportProvider) Start() {
	go r.run()
}

func (r *ReportProvider) Stop() {
	r.stop <- true
}

func (r *ReportProvider) run() {

	for {
		select {
		case reportdata := <-r.report:
			err := r.inform(reportdata)
			if err != nil {
				r.logging.Error(err.Error())
			}
		case <-r.stop:
			fmt.Println("ReportProvider received stop signal")
			return
		}
	}
}

func (r *ReportProvider) transToAbnType(rf ReportForm) ([]byte, error) {

	ra := ReportAbnWithrawalForm{rf.Coin, rf.TxId, rf.Fromaccount, rf.Toaccount}
	byte, err := json.Marshal(ra)

	return byte, err
}

func (r *ReportProvider) transToDepositType(rf ReportForm) ([]byte, error) {

	symbols := []string{}
	symbols = append(symbols, rf.Coin)

	addr := AddressData{rf.Toaccount, symbols}
	address := []AddressData{}
	address = append(address, addr)

	bd := Blockdata{rf.TiptSetNum, rf.TiptSetNum}

	var rd = ReportDepositForm{bd, address}

	byte, err := json.Marshal(rd)

	return byte, err
}

func (r *ReportProvider) inform(rf ReportForm) error {

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

	r.logging.Infof("inform : To: %n \n", rf.Toaccount)

	var url string
	var bytedata []byte

	switch rf.Kind {
	case ABNWITHRWAL_REPORT:
		bytedata, err = r.transToAbnType(rf)
		url = r.abnWUrl

	case DEPOSIT_REPORT:
		bytedata, err = r.transToDepositType(rf)
		url = r.depositUrl
	}

	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytedata))

	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return err
}
