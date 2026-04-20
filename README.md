Ask - AI Powered Terminal Helper

ask is a minimalistic, insanely fast CLI tool that uses AI (Groq/Llama 3) to convert natural language queries into executable terminal commands. It has an automated setup wizard triggered upon the first query. It is designed specifically for Linux.
Full Linux Setup

This guide covers everything you need to know to use this powerful tool.
1. Install Go (The Compiler)

If you don't have Go installed, run the following based on your distribution:

Arch Linux:
Bash

sudo pacman -S go

Ubuntu/Debian:
Bash

sudo apt update && sudo apt install golang

Verify the installation:
Bash

go version

2. Clone and Build

Clone the repository and compile the binary:
Bash

git clone https://github.com/itzAvhi/ask.git
cd ask
go build -o ask main.go

3. Make it Global

To run ask from any folder without typing the full path, move it to your system's binary folder:
Bash

sudo mv ask /usr/local/bin/