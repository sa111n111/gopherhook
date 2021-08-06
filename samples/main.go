package main

import (
	"github.com/sal1entx86/gopherhook"
)

func main() {
	goReq := gopherhook.Webhook{
		Content:  "TEST!!!",
		Username: "WEBHOOKB",
		TTS:      true,
	}

	goReq.Execute("url")
}
