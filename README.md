# Ask - Ai powered Terminal Helper

ask is a minimalistic, insanly fast CLI tool that uses AI (Groq/Llama 3) to convert natural language  querries into executable terminal commands. It has a automated setup wisard triggred upon the first querry of user. It is designed specifically for Linux.

# Full Linux Setup

This guide coveres everything you need to know to use this powerful tool.

# 1. Install Go (The compilar)
If you dont have Go installed, run the following based on you distribution:

#Arch Linux

    ```bash
    sudo pacman -S go
    ```
    
#Ubuntu/Debian:

    ```bash
    sudo apt update && sudo apt install golang
    ```
    
verify the installation:

    ```bash
    go versions
    ```
# 2. Clone and build
Clone the repository and compile the binary:
    ```bash
    git clone https://github.com/itzAvhi/ask.git
    cd ask
    go build -o ask main.go
    ```
# 3. Make it Global
To run ask from any folder without typing the full path, move it to your systems binary folder:

    ```bash
    sudo mv ask /usr/local/bin/
    ```


