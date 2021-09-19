package main

import (
	"github.com/sa111n111/gopherhook"
)

func main() {
	goReq := gopherhook.Webhook{
		Content:  "TEST!!!",
		Username: "WEBHOOKB",
		TTS:      true,
	}

	goReq.Execute("url")
}
