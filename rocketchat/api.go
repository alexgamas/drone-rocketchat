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

// func (rc *RocketChat) DoRequest(req *http.Request) ([]byte, error) {
// 	return rc.request(req)
// }

//PostRequest Low level post request
func (rc *RocketChat) PostRequest(endpoint string, payload interface{}) ([]byte, error) {
	ul := rc.urlApi + endpoint
	body, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, ul, bytes.NewBuffer(body))

	for k, v := range rc.httpHeader() {
		req.Header.Add(k, v)
	}

	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}
	return rc.request(req)
}
