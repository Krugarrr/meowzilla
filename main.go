package meowzilla

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5613559808:AAEcXft_FdEFgIgxAvOePob7GST6_smWjLQ")
	if err != nil {
		log.Panic(err)
	}

	bot.RemoveWebhook()
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Secret button 1"),
				tgbotapi.NewKeyboardButton("Secret button 2"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Open sub-menu"),
			),
		)

		switch update.Message.Text {
		case "/start":
			msg.Text = "Hello! I'm a simple bot for VK task. Press one of the buttons below to get started."
		case "Secret button 1":
			msg.Text = "Meow!"
		case "Secret button 2":
			msg.Text = "Meeeeeeooooooow"
		case "Open sub-menu":
			msg.Text = "Sub-menu opened"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Sub-button 1"),
					tgbotapi.NewKeyboardButton("Sub-button 2"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Close sub-menu"),
				),
			)
		case "Sub-button 1":
			msg.Text = "You pressed Sub-button 1"
		case "Sub-button 2":
			msg.Text = "You pressed Sub-button 2"
		case "Close sub-menu":
			msg.Text = "Sub-menu closed"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Secret button 1"),
					tgbotapi.NewKeyboardButton("Secret button 2"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Open sub-menu"),
				),
			)
		default:
			msg.Text = "I don't understand"
		}

		bot.Send(msg)
	}
}
