/*
2021-02-10

Written by wowlsh93
*/

package eth

import (
	"bytes"
	"encoding/json"
	"github.com/wowlsh93/filecoinscanner/common/flogging"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/crawler/net"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/util"
	"net/http"
)

type EthRPC struct {
	httprpc *net.HTTPRpc
}

func New(url string, options ...func(rpc *EthRPC)) *EthRPC {

	rpc := &EthRPC{
		httprpc: &net.HTTPRpc{
			url,
			http.DefaultClient,
			flogging.GetLogger(),
		},
	}

	for _, option := range options {
		option(rpc)
	}

	return rpc
}

func (e *EthRPC) getBlock(method string, withTransactions bool, params ...interface{}) (*Block, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var response proxyBlock
	if withTransactions {
		response = new(proxyBlockWithTransactions)
	} else {
		response = new(proxyBlockWithoutTransactions)
	}

	err = json.Unmarshal(result, response)
	if err != nil {
		return nil, err
	}

	block := response.toBlock()
	return &block, nil
}

func (e *EthRPC) EthGetBlockByNumber(number int, withTransactions bool) (*Block, error) {
	return e.getBlock("eth_getBlockByNumber", withTransactions, util.IntToHex(number), withTransactions)
}

func (e *EthRPC) EthLastBlockNumber() (int, error) {
	var response string
	if err := e.httprpc.CallWithResult("eth_blockNumber", &response); err != nil {
		return 0, err
	}

	return util.ParseInt(response)
}
