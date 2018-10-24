package rocketchat

type WebHookPostPayload struct {
	Text        string        `json:"text,omitempty"`
	Channel     string        `json:"channel,omitempty"`
	Username    string        `json:"alias,omitempty"`
	IconUrl     string        `json:"avatar,omitempty"`
	IconEmoji   string        `json:"emoji,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

type ChatPostMessageAPIRequest struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
	Avatar  string `json:"avatar"`
}

func (rc *RocketChat) PostMessage(payload *WebHookPostPayload) error {

	_, err := rc.PostRequest(chatPostMessageApiEndpoint, payload)

	if err != nil {
		return err
	}

	return nil

}
