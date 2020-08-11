package main

import (
	"strings"
	"os"
	"log"

	"github.com/gempir/go-twitch-irc"
	"github.com/joho/godotenv"
)

func main() {
    envErr := godotenv.Load()
    if envErr != nil {
        log.Fatal("Error loading .env file")
    }

    twitchUser := os.Getenv("TWITCH_USER")
    twitchOAuth := os.Getenv("TWITCH_OAUTH")
    twitchChannel := os.Getenv("TWITCH_CHANNEL")

	// Bot name and oauth token
	client := twitch.NewClient(twitchUser, twitchOAuth)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		if strings.TrimRight(message.Message, "\n") == "!bot" {
			client.Say(message.Channel, "Hello there "+message.User.Name)
		}

	})

	// User's channel
	client.Join(twitchChannel)

	twitchErr := client.Connect()
	if twitchErr != nil {
		panic(twitchErr)
	}

}
