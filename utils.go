package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"tg-currency-bot/coinranking"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getenv(name string) (variable string) {
	if variable = os.Getenv(name); variable == "" {
		log.Fatalf("Environment variable %s is empty\n", name)
	}

	return variable
}

func newCoinMessage(coin *coinranking.Coin, chatID int64) tgbotapi.Chattable {
	price, _ := strconv.ParseFloat(coin.Price, 32)
	iconURL := "https://raw.githubusercontent.com/aandrosov0/crypto-icons/master/" + coin.Symbol + ".png"

	message := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(iconURL))
	message.Caption = fmt.Sprintf("_%s_ - текущий курс *%.3f$* 💰\nСайт - %s 💳",
		coin.Name, price, coin.WebsiteURL)
	message.ParseMode = tgbotapi.ModeMarkdown

	return message
}
