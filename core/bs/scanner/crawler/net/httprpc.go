/*
2021-02-10

Written by wowlsh93
*/

package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wowlsh93/filecoinscanner/common/flogging"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

type httpClient interface {
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err Error) Error() string {
	return fmt.Sprintf("Error %d (%s)", err.Code, err.Message)
}

type Response struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *Error          `json:"error"`
}

type Request struct {
	ID      int           `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type HTTPRpc struct {
	Url    string
	Client httpClient
	Log    *logrus.Logger
}

func New(url string, options ...func(rpc *HTTPRpc)) *HTTPRpc {
	rpc := &HTTPRpc{
		Url:    url,
		Client: http.DefaultClient,
		Log:    flogging.GetLogger(),
	}
	for _, option := range options {
		option(rpc)
	}

	return rpc
}

func (rpc *HTTPRpc) URL() string {
	return rpc.Url
}

func (rpc *HTTPRpc) CallWithResult(method string, target interface{}, params ...interface{}) error {
	result, err := rpc.Call(method, params...)
	if err != nil {
		return err
	}

	if target == nil {
		return nil
	}

	return json.Unmarshal(result, target)
}

func (rpc *HTTPRpc) Call(method string, params ...interface{}) (json.RawMessage, error) {
	request := Request{
		ID:      1,
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := rpc.Client.Post(rpc.Url, "application/json", bytes.NewBuffer(body))
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, *resp.Error
	}

	return resp.Result, nil

}
