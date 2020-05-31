package cmd

import (
	"fmt"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

func version() {
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

	fmt.Printf("Version: 1.0.0")
	
}
