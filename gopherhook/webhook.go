package gopherhook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Webhook holds the basic structure of a webhook.
// This was written only for the execute method of the discord webhook api.
// Another thing to note is that you will not submit your webhook token,
// we use the raw url of the webhook to send the request.
type Webhook struct {
	// Content is the message contents you'd like the webhook to send.
	Content string

	// Username Overrides the default username of the webhook
	Username string

	// AvatarUrl Overrides the default avatar of the webhook
	AvatarUrl string

	// TTS If the message is text-to-speech or not.
	TTS bool
}

// NewWebhook creates a new webhook object.
// Takes in the same arguments as the Webhook struct.
func (w *Webhook) NewWebhook(content string, username string, avatar_url string, tts bool) Webhook {
	return Webhook{
		Content:   content,
		Username:  username,
		AvatarUrl: avatar_url,
		TTS:       tts,
	}
}

// Execute sends a message in a channel from a webhook.
// Takes in the raw `URL` of the webhook for convenience.
// This assumes you've already provided the fields needed in the webhook struct.
func (w *Webhook) Execute(url string) {
	// The json body. This is to be encoded into json, and then be stored
	// into the newly allocated byte buffer.
	json_body, buff := map[string]interface{}{"content": w.Content, "username": w.Username, "avatar_url": w.AvatarUrl, "tts": w.TTS}, new(bytes.Buffer)

	// Encode the body
	json.NewEncoder(buff).Encode(json_body)

	// Send the stored json data to the URL as a post request.
	http.Post(url, "application/json", buff)
}
