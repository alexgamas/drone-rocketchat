package rocketchat

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//LoginRequest -> https://rocket.chat/docs/developer-guides/rest-api/authentication/login/#payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//LoginResponse -> https://rocket.chat/docs/developer-guides/rest-api/authentication/login/#result
type LoginResponse struct {
	Status string             `json:"status"`
	Data   LoginTokenResponse `json:"data"`
}

//LoginTokenResponse -> https://rocket.chat/docs/developer-guides/rest-api/authentication/login/#result
type LoginTokenResponse struct {
	UserID    string `json:"userId"`
	AuthToken string `json:"authToken"`
}

//Login -> Authenticate with the REST API. [/api/v1/login]
//https://rocket.chat/docs/developer-guides/rest-api/authentication/login/
func (rc *RocketChat) Login(payload *LoginRequest) error {

	//body io.Reader
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(rc.urlApi+loginApiEndPoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	t, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return errors.New(string(t))
	}

	res := new(LoginResponse)
	err = json.Unmarshal(t, res)
	if err != nil {
		return err
	}

	rc.userId = res.Data.UserID
	rc.authToken = res.Data.AuthToken

	return nil
}
