package telegram

import (
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tbauriedel/go-telegram-icinga2/internal"
	"log"
)

type Telegram struct {
	BotToken string
	ChatID   int64
}

// New returns new instance of Telegram
func New(token string, chatid int64) (telegram *Telegram) {
	telegram = &Telegram{
		BotToken: token,
		ChatID:   chatid,
	}

	return
}

// SendMessage sends given message
func (telegram *Telegram) SendMessage(message string) {
	bot, err := telegramBotApi.NewBotAPI(internal.TelegramBotToken)
	if err != nil {
		log.Fatalf("cant create api instance: %s", err)
	}
	bot.Debug = true

	msg := telegramBotApi.NewMessage(telegram.ChatID, message)
	if _, err = bot.Send(msg); err != nil {
		log.Fatalf("cant send message: %s", err)
	}
}
