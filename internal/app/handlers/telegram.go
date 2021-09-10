package handlers

import (
	"Project/config"
	"Project/internal/app/services"
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramHandler struct {
	balanceService *services.BalanceService
	bot            *tgbotapi.BotAPI
}

func NewTelegramHandler(balanceService *services.BalanceService) *TelegramHandler {
	return &TelegramHandler{
		balanceService: balanceService,
	}
}

func (t *TelegramHandler) Telebot() {
	var err error

	t.bot, err = tgbotapi.NewBotAPI(config.PgvTestBotToken)
	if err != nil {
		log.Panic(err)
	}

	t.bot.Debug = true

	log.Printf("Authorized on account %s", t.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := t.bot.GetUpdatesChan(u)

	for update := range updates {
		t.handleUpdates(update)
	}
}

func (t *TelegramHandler) handleUpdates(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if update.Message.Text == "/getBalance" {
		t.handleBalanceMessage(update)
	} else {
		// ПИСАТЬ ТУТ
	}
	// if update.Message.Text == "/getBalance2" {
	// 	t.handleBalance2Message(update)
	// }
}

func (t *TelegramHandler) handleBalanceMessage(update tgbotapi.Update) {
	balance, err := t.balanceService.GetBalance()
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка")
		t.bot.Send(msg)
	}

	bytes, err := json.Marshal(balance)
	if err != nil {
		log.Println(err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(bytes))

	t.bot.Send(msg)
}
