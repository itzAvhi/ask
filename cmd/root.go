package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Ask",
	Short: "AI powered terminal helper",
	Long:  "AI powered terminal helper that accompanies the user by finding the appropiate Terminal command as per the users requirement",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi, I am 'Ask', How can i make your terminal session easy? ")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
