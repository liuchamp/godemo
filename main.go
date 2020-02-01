package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("821401577:AAEJ5D6O8evdFul6bd4pc1hBuSQsI8jfzbA")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	chat, err := bot.GetChat(tgbotapi.ChatConfig{
		ChatID:             -375956450,
		SuperGroupUsername: "yexijgyun",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println("is channel --->", chat.IsChannel())
	log.Println("is group --->", chat.IsGroup())
	log.Println("is private --->", chat.IsPrivate())
	log.Println("is supper group --->", chat.IsSuperGroup())
	log.Println("is --->", chat.Type, chat.UserName, chat.Title)

	rmsg := tgbotapi.NewMessage(-375956450, "bot is ok")
	if _, err := bot.Send(rmsg); err != nil {
		log.Panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we should leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		println("to------>", update.Message.Chat.ID)
		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "<pre>|  命令   |    说明      |      示例        | \n"
			msg.Text += fmt.Sprintf("|%-8s|%-12s|%-16s|\n", "/stauts", "获取服务状态，默认是全部服务", "/status webserver")
			msg.Text += fmt.Sprintf("|%-8s|%-12s|%-16s|\n", "/log", "获取 log 服务器的状态", "/log")
			msg.Text += "</pre>"
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		case "open":
			msg.ReplyMarkup = numericKeyboard
			msg.Text = update.Message.Text
		case "close":
			msg.Text = update.Message.Text
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
		default:
			msg.Text = "I don't know that command"
		}
		msg.ParseMode = tgbotapi.ModeHTML
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
