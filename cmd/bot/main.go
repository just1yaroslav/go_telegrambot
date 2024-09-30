package main

import (
	"flag"
	"github.com/just1yaroslav/go_telegrambot/clients/telegram"
	"log"
)

// Так же можно получать из флага
const (
	tgBotHost = "api.telegram.org"
)

func main() {
	// tgCleint
	tgClient := telegram.New(tgBotHost, mustToken())

	// fetcher

	// processor

	// consumer.start
}

func mustToken() (string, error) {
	token := flag.String(
		"8128506763:AAHg0pxw5tprVXCnznRkBBU6uuZKmqrhLmM",
		"",
		"token for telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("bot no have token-key")
	}

}
