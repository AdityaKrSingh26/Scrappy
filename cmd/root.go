package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/AdityaKrSingh26/Scrappy/config"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scrappy",
	Short: "Scrappy is a CLI tool for scraping internship opportunities.",
	Long: `Scrappy is a command-line tool built in Golang for scraping and managing
internship listings from various websites. It offers commands to run scraping 
tasks on demand or at regular intervals using a cron job.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This function is called by main.main() and needs to happen only once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Define global flags here (if needed)
	// For example:
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scrappy.yaml)")
}

// initConfig reads in config file and environment variables if set.
func initConfig() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}
	// Optionally read environment variables that match.
	viper.AutomaticEnv()
}
