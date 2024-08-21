package main

import (
	"flag"
	"log"
	"telegram-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	t := MustToken()
	tgClient := telegram.New(tgBotHost, t)
	_ = tgClient

}

func MustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
