package cmd

import (
	"fmt"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

func botsetup() {
	botAPIToken := ""
	if viper.GetString("BotAPIToken") != "" {
		botAPIToken = viper.GetString("BotAPIToken")
	} else {
		fmt.Println("'BotAPIToken' is not configured. Maybe configfile is missing.")
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(botAPIToken)
	if err != nil {
		fmt.Println("Connection to Telegram API failed:")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)
	fmt.Printf("You could open now Telegam, connect to your bot https://t.me/%s and type /getchatid to get your chatid.", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

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
		chatid := strconv.FormatInt(update.Message.Chat.ID, 10)

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "type /getchatid to get your chatid\n /version to get your version\nPlease see instructions on <a href='https://github.com/joshuabeny1999/telegram-server-update-bot/blob/master/Readme.md'>Github</a> for more details"
		case "getchatid":
			msg.Text = "Your chatid is: <pre>" + chatid + "</pre>"
		case "version":
			msg.Text = "Version: 1.0.0"
		default:
			msg.Text = "type /getchatid to get your chatid\nPlease see instructions on <a href='https://github.com/joshuabeny1999/telegram-server-update-bot/blob/master/Readme.md'>Github</a> for more details"
		}
		msg.ParseMode = "HTML"

		if _, err := bot.Send(msg); err != nil {
			fmt.Println("Could not send message.")
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
