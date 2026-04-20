package cmd

import (
	"context"
<<<<<<< HEAD
	"fmt"
	"os"
	"os/exec"
=======
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
>>>>>>> 70ebb87 (Added better linux intregations)
	"strings"

	"github.com/fatih/color"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

<<<<<<< HEAD
=======
type Config struct {
	APIKey string `json:"api_key"`
	OS     string `json:"os"`
}

>>>>>>> 70ebb87 (Added better linux intregations)
var (
	name    string
	confirm string
)

<<<<<<< HEAD
var rootCmd = &cobra.Command{
	Use:   "ask [query]",
	Short: "AI powered terminal helper",
	Long:  "AI powered terminal helper that finds the appropriate terminal command for your requirements.",
	// This allows the root command to take your natural language query as arguments
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Combine all arguments into one string
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

		// 2. Configure client for Groq
		config := openai.DefaultConfig(apiKey)
=======
func getConfigPath() string {
	configDir, _ := os.UserConfigDir()
	return filepath.Join(configDir, "ask", "config.json")
}

func runSetup() Config {
	var key string
	fmt.Print("Enter your api key: ")
	fmt.Scanln(&key)

	cfg := Config{
		APIKey: strings.TrimSpace(key),
		OS:     runtime.GOOS,
	}

	path := getConfigPath()
	os.MkdirAll(filepath.Dir(path), 0700)
	data, _ := json.MarshalIndent(cfg, "", "  ")
	os.WriteFile(path, data, 0600)

	return cfg
}

var rootCmd = &cobra.Command{
	Use:   "ask [query]",
	Short: "AI powered terminal helper",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		path := getConfigPath()
		var cfg Config

		if _, err := os.Stat(path); os.IsNotExist(err) {
			cfg = runSetup()
		} else {
			data, _ := os.ReadFile(path)
			json.Unmarshal(data, &cfg)
		}

		query := strings.Join(args, " ")
		if query == "" {
			fmt.Println("Usage: ask <query>")
			return
		}

		config := openai.DefaultConfig(cfg.APIKey)
>>>>>>> 70ebb87 (Added better linux intregations)
		config.BaseURL = "https://api.groq.com/openai/v1"
		client := openai.NewClientWithConfig(config)

		color.Blue("Thinking...")

<<<<<<< HEAD
		// 3. Request completion from Groq
=======
>>>>>>> 70ebb87 (Added better linux intregations)
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: "llama-3.3-70b-versatile",
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleSystem,
<<<<<<< HEAD
						Content: "You are a Linux CLI expert. Return ONLY the raw command. No markdown, no backticks, no text. Just the command.",
=======
						Content: fmt.Sprintf("You are a %s CLI expert. Return ONLY the raw command. No markdown, no text. If a question is not technical, return 'echo \"Please ask a technical question.\"'", cfg.OS),
>>>>>>> 70ebb87 (Added better linux intregations)
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

<<<<<<< HEAD
		// 4. Sanitize and display the command
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

// execute handles running the shell command safely
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

// Subcommand: Encrypt (Your custom feature)
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts a file",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			color.Red("Error: Please provide a filename with -n")
			return
		}
		fmt.Printf("Encrypting file: %s...\n", name)
		// Add your encryption logic here
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
=======
		command := strings.TrimSpace(resp.Choices[0].Message.Content)
		command = strings.Trim(command, "`")

		fmt.Printf("\nSuggested Command: %s\n\n", command)

		fmt.Print("Run this command? (y/N): ")
		fmt.Scanln(&confirm)
		if strings.ToLower(confirm) == "y" {
			execute(command, cfg.OS)
		}
	},
}

func execute(command string, userOS string) {
	var c *exec.Cmd
	if userOS == "windows" {
		c = exec.Command("powershell", "-Command", command)
	} else {
		c = exec.Command("bash", "-c", command)
	}

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	err := c.Run()
	if err != nil {
		color.Red("Execution failed: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the file")
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts a file",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Println("Error: Please provide a filename with -n")
			return
		}
		fmt.Printf("Encrypting file: %s...\n", name)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
>>>>>>> 70ebb87 (Added better linux intregations)
