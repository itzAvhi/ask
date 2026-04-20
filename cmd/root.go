package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

type Config struct {
	APIKey string `json:"api_key"`
	OS     string `json:"os"`
}

var (
	name    string
	confirm string
)

var rootCmd = &cobra.Command{
	Use:   "ask [query]",
	Short: "AI powered terminal helper",
	Long:  "AI powered terminal helper that finds the appropriate terminal command for your requirements.",

	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		query := strings.Join(args, " ")
		if query == "" {
			color.Yellow("Usage: ask <what you want to do>")
			return
		}

		apiKey := os.Getenv("GROQ_API_KEY")
		if apiKey == "" {
			color.Red("Error: GROQ_API_KEY not set.")
			color.Green("Please run: export GROQ_API_KEY='your_key'")
			return
		}

		config := openai.DefaultConfig(apiKey)
		config.BaseURL = "https://api.groq.com/openai/v1"
		client := openai.NewClientWithConfig(config)

		color.Blue("Thinking...")

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: "llama-3.3-70b-versatile",
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleSystem,
						Content: "You are a Linux CLI expert. Return ONLY the raw command. No markdown, no backticks, no text. Just the command.",
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: query,
					},
				},
			},
		)

		if err != nil {
			color.Red("Error: %v", err)
			return
		}

		command := strings.TrimSpace(resp.Choices[0].Message.Content)
		command = strings.Trim(command, "`")

		color.Yellow("\nSuggested Command: ")
		fmt.Printf("  %s\n\n", command)

		// 5. Confirmation and Execution
		fmt.Print("Run this command? (y/N): ")
		fmt.Scanln(&confirm)
		if strings.ToLower(confirm) == "y" {
			execute(command)
		}
	},
}

func execute(command string) {
	c := exec.Command("bash", "-c", command)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	err := c.Run()
	if err != nil {
		color.Red("Execution failed: %v", err)
	}
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts a file",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			color.Red("Error: Please provide a filename with -n")
			return
		}
		fmt.Printf("Encrypting file: %s...\n", name)

	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
