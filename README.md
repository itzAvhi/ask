# Ask - AI Powered Terminal Helper 

`ask` is a minimalistic, insanely fast CLI tool that uses AI(Groq/Llama 3) to convert natural language querries into executiable terminal commands. It features an automated wisard triggred upon the first querry and is designed specifically for Linux

---

## Full Linux Setup

This guide covers everything you need to know to install and use this tool globally.

### 1. Install Go (The compiler )
if you dont have Go installed, Run the following command based on your distribution

**Arch Linux**
```bash
    sudo pacman -S go
```

**Ubuntu/Debian**
```bash
    sudo apt update && sudo apt install golang
```

**Verify the installation**
```bash
    go version
```

### 2.Clone and Build
Clone the repository and compile the binary:
```bash
    git clone [https://github.com/itzAvhi/ask.git](https://github.com/itzAvhi/ask.git)
    cd ask
    go build -o ask main.go
```
### Make it Global
to run `ask` from any folder without typing the full path, move it to your system's binary folder
```bash
    sudo mv ask /usr/local/bin/
```
