package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"err"
	"err/errType"
)

var Get = func(uv *url.Values, ul string, res interface{}) {
	body := fetch(uv, ul)
	e := json.Unmarshal(body, &res)
	err.Found(e, errType.JSONParse)
}

func fetch(uv *url.Values, ul string) []byte {
	req, e := http.NewRequest(http.MethodGet, ul, nil)
	err.Found(e, errType.HTTPRequest)
	req.URL.RawQuery = (*uv).Encode()
	return request(req)
}

func request(req *http.Request) (content []byte) {
	var client http.Client
	res, e := client.Do(req)
	err.Found(e, errType.CreateClient)
	defer res.Body.Close()

	content, e = ioutil.ReadAll(res.Body)
	err.Found(e, errType.IORead)
	return
}
