/*
2021-02-10

Written by wowlsh93
*/

package fil

import (
	"bytes"
	"encoding/json"
	"github.com/wowlsh93/filecoinscanner/common/flogging"
	"github.com/wowlsh93/filecoinscanner/core/bs/scanner/crawler/net"
	"net/http"
)

type FilRPC struct {
	httprpc *net.HTTPRpc
}

func New(url string, options ...func(rpc *FilRPC)) *FilRPC {

	rpc := &FilRPC{
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

func (e *FilRPC) getTipset(method string, params ...interface{}) (*Tipset, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var tipset Tipset

	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return nil, err
	}

	return &tipset, nil
}

func (e *FilRPC) getMessages(method string, params ...interface{}) (Messages, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return Messages{}, err
	}
	if bytes.Equal(result, []byte("null")) {
		return Messages{}, nil
	}

	var msgs Messages
	err = json.Unmarshal(result, &msgs)
	if err != nil {
		return Messages{}, err
	}

	return msgs, nil
}

func (e *FilRPC) FilGetTipsetByHeight(number int) (*Tipset, error) {
	return e.getTipset("Filecoin.ChainGetTipSetByHeight", number, nil)
}

func (e *FilRPC) FilGetMessagesByCID(blockcid string) (Messages, error) {

	cid := make(map[string]string)
	cid["/"] = blockcid

	msgs, err := e.getMessages("Filecoin.ChainGetBlockMessages", cid)
	msgs.Blockcid = blockcid
	return msgs, err
}
