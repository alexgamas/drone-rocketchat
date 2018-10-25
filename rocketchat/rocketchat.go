package rocketchat

// RocketChat Client
type RocketChat struct {
	urlApi    string
	userId    string
	authToken string
}

//New cria uma api RocketChat
func New(urlAPI string, userID string, authToken string) *RocketChat {
	return &RocketChat{urlAPI, userID, authToken}
}

func (rc *RocketChat) httpHeader() map[string]string {
	h := make(map[string]string)

	h["X-Auth-Token"] = rc.authToken
	h["X-User-Id"] = rc.userId

	return h
}
