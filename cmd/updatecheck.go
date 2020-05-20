package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

func updatecheck() {
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

	out, err := exec.Command("apt", "list", "--upgradeable").Output()
	if err != nil {
		fmt.Println("Load package list failed:")
		fmt.Println(err)
		os.Exit(1)
	}
	packages := strings.Split(string(out), "\n")
	packCount := 0
	message := "<b>" + serverName + "</b>\n"
	packageList := ""
	for _, pack := range packages {
		switch pack {
		case "Listing...":
			continue
		case "":
			continue
		default:
			packSplitted := strings.Split(pack, "/")
			packageList += "- " + packSplitted[0] + "\n"
			packCount++
		}
	}
	if packCount > 0 {
		bot, err := tgbotapi.NewBotAPI(botAPIToken)
		if err != nil {
			fmt.Println("Connection to Telegram API failed:")
			fmt.Println(err)
			os.Exit(1)
		}

		message += "<i>" + strconv.Itoa(packCount) + " packages can be updated:</i>\n\n<pre>" + packageList + "</pre>"
		msg := tgbotapi.NewMessage(chatUserID, message)
		msg.ParseMode = "HTML"

		if _, err := bot.Send(msg); err != nil {
			fmt.Println("Could not send message.")
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
