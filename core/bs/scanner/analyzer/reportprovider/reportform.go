package reportprovider

type KIND string

const (
	DEPOSIT_REPORT     = KIND("deposit")
	ABNWITHRWAL_REPORT = KIND("abnwithdrawal")
)

type ReportForm struct {
	Coin        string
	TiptSetNum  int
	BlockCid    string
	TxId        string
	Kind        KIND // kind : deposit , abnwithdrawal
	Fromaccount string
	Toaccount   string
}

type ReportDepositForm struct {
	Blockdata   Blockdata     `json:"blockData"`
	AddressData []AddressData `json:"addressData"`
}

type Blockdata struct {
	MinBlock int `json:"minBlock"`
	MaxBlock int `json:"maxBlock"`
}
type AddressData struct {
	Address string
	Simbols []string
}

type ReportAbnWithrawalForm struct {
	Symbol      string `json:"symbol"`
	TxId        string `json:"txid"`
	Fromaccount string `json:"from"`
	Toaccount   string `json:"to"`
}
