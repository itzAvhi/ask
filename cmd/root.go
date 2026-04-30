package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var (
	fileName string
	confirm  string
)

func getOSName() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "Linux"
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"")
		}
	}
	return "Linux"
}

var rootCmd = &cobra.Command{
	Use:   "ask [query]",
	Short: "AI powered terminal helper",
	Long:  "A context-aware CLI tool that suggests and executes commands based on your current environment.",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("GROQ_API_KEY")

		if apiKey == "" {
			color.Cyan("Welcome to ask! First-time setup starting...")
			fmt.Print("Please enter your Groq API Key: ")
			var inputKey string
			fmt.Scanln(&inputKey)

			if inputKey == "" {
				color.Red("Error: API Key is required.")
				return
			}

			home, _ := os.UserHomeDir()
			bashrcPath := home + "/.bashrc"
			localBin := home + "/.local/bin"

			os.MkdirAll(localBin, 0755)

			selfPath, _ := os.Executable()
			src, err := os.Open(selfPath)
			if err == nil {
				dst, err := os.OpenFile(localBin+"/ask", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
				if err == nil {
					io.Copy(dst, src)
					dst.Close()
				}
				src.Close()
			}

			configLines := fmt.Sprintf("\n# Added by ask CLI\nexport GROQ_API_KEY='%s'\nexport PATH=\"$PATH:%s\"\n", inputKey, localBin)
			f, err := os.OpenFile(bashrcPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err == nil {
				f.WriteString(configLines)
				f.Close()
				color.Green("✓ API Key saved and 'ask' installed to %s", localBin)
				color.Yellow("Please run 'source ~/.bashrc' to finish.")
			}
			apiKey = inputKey
		}

		query := strings.Join(args, " ")
		if query == "" {
			color.Yellow("Usage: ask <what you want to do>")
			return
		}

		osName := getOSName()
		pwd, _ := os.Getwd()
		lsOutput, _ := executeAndCapture("ls -F")

		memoryFile := "memory.txt"
		var last15 string
		if data, err := os.ReadFile(memoryFile); err == nil {
			lines := strings.Split(string(data), "\n")
			start := len(lines) - 15
			if start < 0 {
				start = 0
			}
			last15 = strings.Join(lines[start:], "\n")
		}

		config := openai.DefaultConfig(apiKey)
		config.BaseURL = "https://api.groq.com/openai/v1"
		client := openai.NewClientWithConfig(config)

		color.Blue("Thinking...")

		systemPrompt := fmt.Sprintf(
			"You are a Linux CLI expert on %s. "+
				"Current Path: %s | Files here: %s. "+
				"Past Interactions: %s. "+
				"Return ONLY the raw bash command. No markdown, no backticks, no explanations.",
			osName, pwd, strings.TrimSpace(lsOutput), last15,
		)

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: "llama-3.3-70b-versatile",
				Messages: []openai.ChatCompletionMessage{
					{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
					{Role: openai.ChatMessageRoleUser, Content: query},
				},
			},
		)

		if err != nil {
			color.Red("API Error: %v", err)
			return
		}

		suggestedCommand := strings.TrimSpace(resp.Choices[0].Message.Content)
		suggestedCommand = strings.Trim(suggestedCommand, "`")

		color.Yellow("\nSuggested Command: ")
		fmt.Printf("  %s\n\n", suggestedCommand)

		fmt.Print("Run this command? (y/N): ")
		fmt.Scanln(&confirm)
		if strings.ToLower(confirm) == "y" {
			execute(suggestedCommand)

			file, err := os.OpenFile(memoryFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err == nil {
				defer file.Close()
				fmt.Fprintf(file, "User: %s\nAsk: %s\n\n", query, suggestedCommand)
			}
		}
	},
}

func execute(command string) {
	if strings.HasPrefix(command, "cd ") {
		target := strings.TrimSpace(strings.TrimPrefix(command, "cd "))
		if strings.HasPrefix(target, "~") {
			home, _ := os.UserHomeDir()
			target = strings.Replace(target, "~", home, 1)
		}
		err := os.Chdir(target)
		if err != nil {
			color.Red("Directory change failed: %v", err)
		} else {
			color.Green("Moved to %s", target)
		}
		return
	}

	c := exec.Command("bash", "-c", command)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	err := c.Run()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return
		}
		color.Red("Execution failed: %v", err)
	}
}

func executeAndCapture(command string) (string, error) {
	var outb bytes.Buffer
	c := exec.Command("bash", "-c", command)
	c.Stdout = &outb
	err := c.Run()
	return outb.String(), err
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts a file (Placeholder)",
	Run: func(cmd *cobra.Command, args []string) {
		if fileName == "" {
			color.Red("Error: Please provide a filename with -n")
			return
		}
		color.Cyan("Encrypting file: %s...", fileName)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringVarP(&fileName, "name", "n", "", "Name of the file to encrypt")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
