package reportprovider

import (
	"encoding/json"
	"testing"
)

func Test_ReportAbnWithrawalForm(t *testing.T) {

	ra := ReportAbnWithrawalForm{"ETH", "0x1234", "0X2222", "Ox3333"}
	byte, _ := json.Marshal(ra)

	if string(byte) != `{"symbol":"ETH","txid":"0x1234","from":"0X2222","to":"Ox3333"}` {
		t.Fatal("Test_ReportAbnWithrawalForm fail ")
	}
}

func Test_ReportDepositForm(t *testing.T) {

	symbols := []string{}
	symbols = append(symbols, "ETH")
	symbols = append(symbols, "BTC")

	addr := AddressData{"0x12", symbols}
	address := []AddressData{}
	address = append(address, addr)

	bd := Blockdata{1, 2}

	var rd = ReportDepositForm{bd, address}

	byte2, _ := json.Marshal(rd)

	var rd2 ReportDepositForm
	json.Unmarshal(byte2, &rd2)

	if rd.Blockdata.MinBlock != rd2.Blockdata.MinBlock {
		t.Fatal("Test_ReportDepositForm fail ")
	}

	if rd.AddressData[0].Address != rd2.AddressData[0].Address {
		t.Fatal("Test_ReportDepositForm fail ")
	}

}
