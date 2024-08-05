package main

import (
	"flag"
)

func main() {
	app := NewApp()

	// Alternatively use flag parsing
	flag.StringVar(&app.WebhookURL, "webhookURL", app.WebhookURL, "Webhook URL")
	flag.StringVar(&app.AvatarUsername, "avatarUsername", app.AvatarUsername, "Avatar username")
	flag.StringVar(&app.AvatarURL, "avatarURL", app.AvatarURL, "Avatar image URL")
	flag.IntVar(&app.ListenPort, "listenPort", app.ListenPort, "Listen port")
	flag.Parse()

	app.Run()
}
