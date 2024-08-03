package main

import (
	"flag"
)

func main() {
	app := NewApp()

	flag.StringVar(&app.WebhookURL, "webhookURL", "https://discord.com/api/webhooks/", "Webhook URL")
	flag.StringVar(&app.AvatarUsername, "avatarUsername", "Annalee Call", "Avatar username")
	flag.StringVar(&app.AvatarURL, "avatarURL", "https://static1.personality-database.com/profile_images/6604632de9954b4d99575e56404bd8b7.png", "Avatar image URL")
	flag.IntVar(&app.ListenPort, "listenPort", 8081, "Listen port")
	flag.Parse()

	app.Run()
}
