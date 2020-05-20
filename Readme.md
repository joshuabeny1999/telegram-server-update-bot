![](https://github.com/joshuabeny1999/telegram-server-update-bot/workflows/Buildw%20Release/badge.svg)

Telegram Server Update Bot
==========================
This is a program which can notify a User via Telegram about Ubuntu Server upgrades. It use commant `apt list --upgradeable` for determine new package updates.

The program is written in [go](golang.org).

(Currently in development)

Installation and Configuration
------------------------------
### 1. Get Telegram Bot API Token
You need to create an telegram Bot to get an API-Key. For creating a Bot do following steps:
1. Open with Telegram bot https://t.me/@BotFather
2. Type `/newbot` and follow instructions to create a bot
3. Now you will recieve an Telegram bot API Token.

### 2. Install Application
1. Download and Upload program to your Ubuntu Server. You find it under [Releases](https://github.com/joshuabeny1999/telegram-server-update-bot/releases/latest) (Check your ARM architecture with `dpkg --print-architecture`) 
2. Untar it: `tar -xvzf telegram-server-update-bot_x.y.z_Linux_amd64.tar.g`
3. Make it executeable: `chmod +x telegram-server-update-bot`
4. Place it on the server in a location you like. For example `/home/<<userdir>>/telegram-server-update-bot`
5. Create an config yaml `/home/<<userdir>>/.telegram-server-update-bot.yaml`
```yaml
BotAPIToken: "YOUR_TELEGRAM_API_TOKEN"
```
### 3. Get your ChatUserID
1. Start program on your Server in setup mode: `./telegram-server-update-bot botsetup`
2. Open now your own Telegram bot with the Telegram-App and send command `/getchatid`
3. You should recieve now the chatid
4. On your Server stop program with keyboard shortcut `CTRL+C`
5. Extend now the config with the chatUserID
```yaml
BotAPIToken: "YOUR_TELEGRAM_API_TOKEN"
ChatUserID: "YOUR_CHAT_USER_ID"
```

### 4. Get notified
To get now notified about new package updates, you could run the script periodically via cron.
For example you could check for updates every day at 6pm (Use `crontab -e ` for adding a cron entry):
```
 0 18 * * * /home/<<userdir>>/telegram-server-update-bot updatecheck 
```

Optional Configuration
-----------------------
### ServerName
It is possible to use the same API Token on multiple Ubuntu servers. So you get notified for more than one server.
To see which server has the update, add following configuration to your yaml file:
```yaml
ServerName: "My Awesome Server"
```

This name will be used as title in each message. If no ServerName is set. It will use as name `Server`.
If you now on each server put an different server name in the config, you can tell them apart insite the Telegram Chat.

### Config File location
You could place the yaml configuration file whereever you want. If it is another location than the home directory,
you could run the program with the option `--config` and provide custom location to config file.
