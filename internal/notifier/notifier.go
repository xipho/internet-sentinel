package notifier

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Notifier interface {
	Notify(message string) error
}

type tgNotifier struct {
	bot    *tgbotapi.BotAPI
	chatId int64
}

func NewNotifier(token string, chatId int64) Notifier {
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	notifier := tgNotifier{
		bot:    botApi,
		chatId: chatId,
	}
	return &notifier
}

func (n *tgNotifier) Notify(message string) error {
	messageConfig := tgbotapi.NewMessage(n.chatId, message)
	_, err := n.bot.Send(messageConfig)
	return err
}
