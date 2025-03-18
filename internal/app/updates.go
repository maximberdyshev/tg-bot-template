package app

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (app *App) message(msg *tgbotapi.Message) {
	switch {
	case msg.IsCommand():
		log.Println("Recieved: new command --", msg.Command())

		if msg.Command() == "test" {
			msg := tgbotapi.NewMessage(msg.Chat.ID, "inline keyboard example")
			msg.ReplyMarkup = defaultIKb

			if _, err := app.Bot.Send(msg); err != nil {
				log.Println("Send message error:", err)
			}
		}

	default:
		log.Println("Recieved: new message --", msg.Text)
	}

}

func (app *App) callback(cb *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := app.Bot.Request(callback); err != nil {
		log.Println("Send request callback error:", err)
	}

	if cb.Data != "back" {
		msg := tgbotapi.NewEditMessageTextAndMarkup(
			cb.Message.Chat.ID,
			cb.Message.MessageID,
			fmt.Sprintf("your choice: %s", cb.Data),
			backwardIKB,
		)

		if _, err := app.Bot.Send(msg); err != nil {
			log.Println("Send message error:", err)
		}
	} else {
		msg := tgbotapi.NewEditMessageTextAndMarkup(
			cb.Message.Chat.ID,
			cb.Message.MessageID,
			"inline keyboard example",
			defaultIKb,
		)

		if _, err := app.Bot.Send(msg); err != nil {
			log.Println("Send message error:", err)
		}
	}

	log.Println("Recieved: new callback data --", cb.Data)
}
