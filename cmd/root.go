package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	goVersion "go.hein.dev/go-version"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "telegram-server-update-bot",
	Short: "Check for server upgrades and send Telegram messages",
	Long:  `Check for Ubuntu server upgrades with apt and then send an Telegram message using a Telegram bot.`,
}

var cmdUpdateCheck = &cobra.Command{
	Use:   "update-check",
	Short: "Check if Server upgrade available and send message",
	Long: `Run this script periodically via cron. 
	    It will check for server updates and send an message via Telegram if any package update is available.`,
	Run: func(cmd *cobra.Command,
		args []string) {
		updatecheck()
	},
}

var cmdBotSetup = &cobra.Command{
	Use:   "botsetup",
	Short: "Telegram Bot setup",
	Long:  `Call programm with this option to setup the bot and get the User-ID`,
	Run: func(cmd *cobra.Command,
		args []string) {
		botsetup()
	},
}

var (
	shortened  = false
	version    = "dev"
	commit     = "none"
	date       = "unknown"
	output     = "json"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version will output the current build information",
		Long:  ``,
		Run: func(_ *cobra.Command, _ []string) {
			resp := goVersion.FuncWithOutput(shortened, version, commit, date, output)
			fmt.Print(resp)
			return
		},
	}
)

// Exec command to run script
func Exec() {
	rootCmd.AddCommand(cmdUpdateCheck)
	rootCmd.AddCommand(cmdBotSetup)
	rootCmd.AddCommand(versionCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	versionCmd.Flags().BoolVarP(&shortened, "short", "s", false, "Print just the version number.")
	versionCmd.Flags().StringVarP(&output, "output", "o", "json", "Output format. One of 'yaml' or 'json'.")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is path-to-script/.telegram-server-update-bot.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".telegram-server-update-bot" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".telegram-server-update-bot")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
