package main

import (
	flag "github.com/spf13/pflag"
	"github.com/tbauriedel/go-telegram-icinga2/internal"
	"github.com/tbauriedel/go-telegram-icinga2/pkg/format"
	"github.com/tbauriedel/go-telegram-icinga2/pkg/telegram"
)

var (
	message string
)

func main() {
	handleArgs()

	if internal.ServiceName != "" {
		message = format.ServiceMessage()
	} else {
		message = format.HostMessage()
	}

	t := telegram.New(internal.TelegramBotToken, internal.TelegramChatId)
	t.SendMessage(message)
}

// handleArgs handles the arguments
func handleArgs() {
	flag.StringVar(&internal.Hostname, "host", "", "Hostname")
	flag.StringVar(&internal.HostAddress, "address", "", "Host address")
	flag.StringVar(&internal.ServiceName, "service", "", "Service name")
	flag.StringVar(&internal.ObjectOutput, "output", "", "Host or service output")
	flag.StringVar(&internal.ObjectState, "state", "", "Host or service state")
	flag.StringVar(&internal.DateTime, "date", "", "Datetime")
	flag.Int64Var(&internal.TelegramChatId, "chatId", 0, "Telegram chat id")
	flag.StringVar(&internal.TelegramBotToken, "token", "", "Telegram bot token")

	flag.CommandLine.SortFlags = false
	flag.Parse()
}
