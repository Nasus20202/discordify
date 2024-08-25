package discord

type StatusRequest struct {
	*CustomStatus `json:"custom_status"`
}

type CustomStatus struct {
	Text      string `json:"text"`
	EmojiName string `json:"emoji_name"`
}

func NewStatusRequest(status string, emoji string) *StatusRequest {
	return &StatusRequest{
		&CustomStatus{
			Text:      status,
			EmojiName: emoji,
		},
	}
}
