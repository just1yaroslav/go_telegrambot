package main

import (
	"flag"
	tgClient "github.com/just1yaroslav/go_telegrambot/clients/telegram"
	"github.com/just1yaroslav/go_telegrambot/consumer/event-consumer"
	"github.com/just1yaroslav/go_telegrambot/events/telegram"
	"github.com/just1yaroslav/go_telegrambot/storage/files"
	"log"
)

// Так же можно получать из флага
const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("server has been started!")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	// Updated call to Start()
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
