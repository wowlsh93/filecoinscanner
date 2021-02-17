/*
2021-02-10

Written by wowlsh93
*/

package fil

type ChainDataResult struct {
	Error        error
	Height       int
	MessagesList []Messages
}

type CID struct {
	Cid string `json:"/"`
}

type Tipset struct {
	Cids []CID `json:"Cids"`
}

type BlsMessages struct {
	Version  int
	To       string
	From     string
	Nonce    int64
	GasPrice string
	GasLimit int64
	Method   int
	Params   string
}

type SecpkMessages struct {
	Version  int
	To       string
	From     string
	Nonce    int64
	GasPrice string
	GasLimit int64
	Method   int
	Params   string
}

type Messages struct {
	Blockcid      string
	BlsMessages   []BlsMessages   `json:"BlsMessages"`
	SecpkMessages []SecpkMessages `json:"SecpkMessages"`
	Cids          []CID           `json:"Cids"`
}
