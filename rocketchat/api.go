package rocketchat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var httpClient *http.Client

func init() {
	httpClient = http.DefaultClient
}

func (rc *RocketChat) request(req *http.Request) ([]byte, error) {
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (rc *RocketChat) DoRequest(req *http.Request) ([]byte, error) {
	return rc.request(req)
}

func (rc *RocketChat) PostRequest(endpoint string, payload interface{}) ([]byte, error) {
	ul := rc.urlApi + endpoint
	body, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, ul, bytes.NewBuffer(body))

	for k, v := range rc.httpHeader() {
		req.Header.Add(k, v)
	}

	if err != nil {
		return nil, err
	}
	return rc.request(req)
}
