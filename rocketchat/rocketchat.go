package rocketchat

type RocketChat struct {
	urlApi    string
	userId    string
	authToken string
}

// cria uma api RocketChat
func New(urlApi string, userId string, authToken string) *RocketChat {
	return &RocketChat{urlApi, userId, authToken}
}

func (rc *RocketChat) httpHeader() map[string]string {
	h := make(map[string]string)

	h["X-Auth-Token"] = rc.authToken
	h["X-User-Id"] = rc.userId

	return h
}
