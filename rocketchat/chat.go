package rocketchat

// ChatPostMessageAPIRequest -> https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage/#payload
type ChatPostMessageAPIRequest struct {
	Text        string        `json:"text,omitempty"`
	Channel     string        `json:"channel,omitempty"`
	Username    string        `json:"alias,omitempty"`
	IconUrl     string        `json:"avatar,omitempty"`
	IconEmoji   string        `json:"emoji,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

//ChatPostMessageAPIResult ..
type ChatPostMessageAPIResult struct {
	Success bool
}

//ChatPinMessageAPIRequest -> https://rocket.chat/docs/developer-guides/rest-api/chat/pinmessage/#payload
type ChatPinMessageAPIRequest struct {
	MessageId string `json:"messageId"`
}

//PostMessage -> Posts a new chat message. [/api/v1/chat.postMessage]
//Payload: https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage/#payload
func (rc *RocketChat) PostMessage(payload *ChatPostMessageAPIRequest) error {

	_, err := rc.PostRequest(chatPostMessageApiEndpoint, payload)

	if err != nil {
		return err
	}

	return nil

}

//PinMessage -> Pins a chat message to the message’s channel. [/api/v1/chat.pinMessage]
//Payload: https://rocket.chat/docs/developer-guides/rest-api/chat/pinmessage/#payload
func (rc *RocketChat) PinMessage(payload *ChatPinMessageAPIRequest) error {

	_, err := rc.PostRequest(chatPinMesssageApiEndpoint, payload)

	if err != nil {
		return err
	}

	return nil
}
