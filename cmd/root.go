package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/AdityaKrSingh26/Scrappy/config"
)

var rootCmd = &cobra.Command{
    Use:   "scrappy",
    Short: "Scrappy is a CLI tool for scraping internship opportunities.",
    Long:  `Scrappy is a command-line tool built in Golang for scraping...`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
}

func initConfig() {
    err := config.LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        os.Exit(1)
    }
    viper.AutomaticEnv()
}
