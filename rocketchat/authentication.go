package rocketchat

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status string             `json:"status"`
	Data   LoginTokenResponse `json:"data"`
}

type LoginTokenResponse struct {
	UserID    string `json:"userId"`
	AuthToken string `json:"authToken"`
}

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
