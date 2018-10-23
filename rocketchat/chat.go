package rocketchat

type ChatPostMessageAPIRequest struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
	Avatar  string `json:"avatar"`
}

func (rc *RocketChat) ChatPostMessage(channelId string, text string) error {

	req := ChatPostMessageAPIRequest{Channel: "geral", Text: "Teste de mensagem", Avatar: "http://res.guggy.com/logo_128.png"}

	_, err := rc.PostRequest(chatPostMessageApiEndpoint, req)

	if err != nil {
		return err
	}

	return nil
}
