package rocketchat

type ChatPostMessageAPIRequest struct {
	Text        string        `json:"text,omitempty"`
	Channel     string        `json:"channel,omitempty"`
	Username    string        `json:"alias,omitempty"`
	IconUrl     string        `json:"avatar,omitempty"`
	IconEmoji   string        `json:"emoji,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

type ChatPostMessageAPIResult struct {
	Success bool
}

type ChatPinMessageAPIRequest struct {
	MessageId string `json:"messageId"`
}

func (rc *RocketChat) PostMessage(payload *ChatPostMessageAPIRequest) error {

	_, err := rc.PostRequest(chatPostMessageApiEndpoint, payload)

	if err != nil {
		return err
	}

	return nil

}

func (rc *RocketChat) PinMessage(payload *ChatPinMessageAPIRequest) error {

	_, err := rc.PostRequest(chatPinMesssageApiEndpoint, payload)

	if err != nil {
		return err
	}

	return nil
}
