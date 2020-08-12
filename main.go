package main

import (
	"log"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load in the .env file
	env := godotenv.Load()

	// If there's an error loading it in, log it
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	// Twitch deets - user and oauth should be your bots deets. Channel is your personal channel/username
	twitchUser := os.Getenv("TWITCH_USER")
	twitchOAuth := os.Getenv("TWITCH_OAUTH")
	twitchChannel := os.Getenv("TWITCH_CHANNEL")

	// Set bot name and oauth token
	client := twitch.NewClient(twitchUser, twitchOAuth)

	// Listen out for new messages
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		// If someone types !bot, reply Hello there *username*
		if strings.TrimRight(message.Message, "\n") == "!bot" {
			client.Say(message.Channel, "Hello there "+message.User.Name)
		}

	})

	// Bot joins user's channel
	client.Join(twitchChannel)

	// Connect to Twitch client
	err := client.Connect()

	// If twitch errors, panic!
	if err != nil {
		panic(err)
	}

}
