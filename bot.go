package main

import (
	"log"
	"tg-currency-bot/coinranking"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var currencyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("BTC"),
		tgbotapi.NewKeyboardButton("ETH"),
		tgbotapi.NewKeyboardButton("TON"),
	),
)

type Bot struct {
	tgApi   *tgbotapi.BotAPI
	coinAPI *coinranking.Client
}

func (b *Bot) Start(tgToken, coinrankingToken string) {
	var err error

	if b.tgApi, err = tgbotapi.NewBotAPI(tgToken); err != nil {
		log.Fatalln("Cannot create bot: " + err.Error())
	}

	b.tgApi.Debug = true
	b.coinAPI = coinranking.NewClient(coinrankingToken)

	updateCfg := tgbotapi.NewUpdate(0)
	updateCfg.Timeout = 60
	b.update(updateCfg)
}

func (b *Bot) update(updateCfg tgbotapi.UpdateConfig) {
	for update := range b.tgApi.GetUpdatesChan(updateCfg) {
		if update.Message != nil {
			message := b.command(update.Message.Text, update.Message.Chat.ID)
			b.tgApi.Send(message)
		}
	}
}

func (b *Bot) command(command string, chatID int64) tgbotapi.Chattable {
	var err error
	var coin *coinranking.Coin

	message := tgbotapi.NewMessage(chatID, "")

	switch command {
	case "/start":
		message.Text = "Привет! Выбери валюту, чтобы узнать её курс 🌍"
		message.ReplyMarkup = currencyKeyboard
	case "BTC":
		coin, err = b.coinAPI.GetCoin(coinranking.BTC)
	case "ETH":
		coin, err = b.coinAPI.GetCoin(coinranking.ETH)
	case "TON":
		coin, err = b.coinAPI.GetCoin(coinranking.TON)
	}

	log.Println(err)

	if coin != nil {
		return newCoinMessage(coin, chatID)
	}

	return message
}
