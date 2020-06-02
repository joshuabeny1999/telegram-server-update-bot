package cmd

import (
	"fmt"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

func testmessage(testmsg string) {
	botAPIToken := ""
	if viper.GetString("BotAPIToken") != "" {
		botAPIToken = viper.GetString("BotAPIToken")
	} else {
		fmt.Println("'BotAPIToken' is not configured. Maybe configfile is missing.")
		os.Exit(1)
	}

	var chatUserID int64
	if viper.GetString("ChatUserID") != "" {
		chatUserID, _ = strconv.ParseInt(viper.GetString("ChatUserID"), 10, 64)
	} else {
		fmt.Println("'ChatUserID' is not configured. Maybe configfile is missing.")
		os.Exit(1)
	}

	serverName := "Server"
	if viper.GetString("ServerName") != "" {
		serverName = viper.GetString("ServerName")
	}

	if testmsg == "" {
		testmsg = "Hello from the telegram-server-update-bot!"
	}

	bot, err := tgbotapi.NewBotAPI(botAPIToken)
	if err != nil {
		fmt.Println("Connection to Telegram API failed:")
		fmt.Println(err)
		os.Exit(1)
	}

	message := "<b>" + serverName + "</b>\n" + testmsg
	msg := tgbotapi.NewMessage(chatUserID, message)
	msg.ParseMode = "HTML"

	if _, err := bot.Send(msg); err != nil {
		fmt.Println("Could not send message.")
		fmt.Println(err)
		os.Exit(1)
	}
}
