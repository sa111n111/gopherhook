package main

import (
	"gopherhook/gopherhook"
)

func main() {
	goReq := gopherhook.Webhook{
		Content:  "TEST!!!",
		Username: "WEBHOOKB",
		TTS:      true,
	}

	goReq.Execute("url")
}
