package rocketchat

//AttachmentField -> https://rocket.chat/docs/developer-guides/rest-api/chat/sendmessage/#attachment-field-objects
type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

//Attachment -> https://rocket.chat/docs/developer-guides/rest-api/chat/sendmessage/#attachments-detail
type Attachment struct {
	Color string `json:"color,omitempty"`

	AuthorName string `json:"author_name,omitempty"`
	AuthorLink string `json:"author_link,omitempty"`
	AuthorIcon string `json:"author_icon,omitempty"`

	Title     string `json:"title,omitempty"`
	TitleLink string `json:"title_link,omitempty"`
	Text      string `json:"text"`

	ImageURL string `json:"image_url,omitempty"`
	ThumbURL string `json:"thumb_url,omitempty"`

	TimeStamp int64 `json:"ts,omitempty"`

	Fields []*AttachmentField `json:"fields,omitempty"`
}
